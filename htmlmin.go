package tinyhtml

import (
	"io"
)

//Made to wrap around a file reader to compress html for webserver applications in order to reduce bandwidth
type Minimizer struct {
	inp     io.Reader
	minflag bool
	buffer  *Queue
	comment bool
	intag   bool
	intext  bool
}

//Creates an html minimizer with the given Reader as its input
func NewMinimizer(i io.Reader) *Minimizer {
	m := new(Minimizer)
	m.inp = i
	m.minflag = true
	m.buffer = new(Queue)
	return m
}

//Read compressed html into the buffer given
func (m *Minimizer) Read(b []byte) (int, error) {
	ob := make([]byte, 1)
	i := 0
	var sb byte
	for i < len(b) {
		if m.buffer.Size() > 0 {
			sb = m.buffer.Pop()
		} else {
			rn, err := m.inp.Read(ob)
			if rn != 1 {
				//what happened?
			}
			if err != nil {
				return i, err
			}
			sb = ob[0]
		}
		switch sb {
		case '-':
			if m.comment {
				temp := make([]byte, 2)
				m.inp.Read(temp)
				if string(temp) == "->" {
					m.comment = false
				}
				continue
			}
		case '<':
			if m.comment {
				continue
			}
			temp := make([]byte, 3)
			m.inp.Read(temp)
			if string(temp) == "!--" {
				m.comment = true
				continue
			} else {
				m.buffer.PushMany(temp)
			}
			m.intag = true
			m.intext = false
		case '>':
			if m.comment {
				continue
			}
			m.intag = false
		case '\n':
			continue
		case '\r':
			continue
		case '\t':
			continue
		case ' ':
			if !m.intext && !m.intag {
				continue
			}
		default:
			if !m.intag {
				m.intext = true
			}
		}
		if !m.comment {
			b[i] = sb
			i++
		}
	}
	return i, nil
}
