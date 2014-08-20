package main

import(
	"code.google.com/p/liblundis/lmath/expr"
	"strings"
	"strconv"
	"errors"
	"math"
	"fmt"
)

func (self *ApproxGUI) interpretImageDim() error {
	strx := strings.Replace(self.dimx_box.Text(), " ", "", -1)
	stry := strings.Replace(self.dimy_box.Text(), " ", "", -1)
	x, err := strconv.ParseInt(strx, 10, 32)
	if err != nil {
		self.err = err
		return err
	}
	y, err := strconv.ParseInt(stry, 10, 32)
	if err != nil {
		self.err = err
		return err
	}
	if x < 100 || y < 100 || x > 2000 || y > 2000 {
		self.err = errors.New("Graph size must be between 100x100px and 2000x2000px")
		return self.err
	}
	self.dimx = int(x)
	self.dimy = int(y)
	return nil
}

func (self *ApproxGUI) interpretSettings() error {
	self.err = self.interpretDegrees()
	if self.err != nil {
		return self.err
	}
	self.err = self.interpretIntervals()
	if self.err != nil {
		return self.err
	}
	self.err = self.interpretAccuracy()
	if self.err != nil {
		return self.err
	}
	self.err = self.interpretFunction()
	if self.err != nil {
		return self.err
	}
	return nil
}

func (self *ApproxGUI) interpretDegrees() error {
	text := strings.Replace(self.degrees_box.Text(), " ", "", -1)
	split := strings.Split(text, ",")
	self.degrees = make([]int, 0, 10)
	for _, str := range split {
		err := self.interpretDegree(str)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *ApproxGUI) interpretDegree(num string) error {
	if hyph := strings.Index(num, "-"); hyph != -1 {
		first, err1 := strconv.ParseInt(num[:hyph], 10, 16)
		if err1 != nil {
			return err1
		}
		second, err2 := strconv.ParseInt(num[hyph+1:], 10, 16)
		if err2 != nil {
			return err2
		}
		if first >= second || first < 0 {
			return errors.New(fmt.Sprintf("first value (%v) in range must be >= 0 and < second(%v)", first, second))
		} else {
			for i := int(first); i <= int(second); i++ {
				self.degrees = append(self.degrees, i)
			}
		}
	} else {
		int_, err := strconv.ParseInt(num, 10, 16)
		if err != nil {
			return err
		}
		self.degrees = append(self.degrees, int(int_))
	}
	return nil
}

func (self *ApproxGUI) interpretIntervals() error {
	str := strings.Replace(self.interval_box.Text(), " ", "", -1)
	split := strings.Split(str, ",")
	if len(split) != 2 {
		return errors.New("interval endpoints must be separated by a comma")
	}
	first, err := expr.EvalCommonf64(split[0])
	if err != nil {
		return err
	}
	second, err := expr.EvalCommonf64(split[1])
	if err != nil {
		return err
	}
	self.ival_start = first
	self.ival_end = second
	return nil
}

// checks the function dropdown menu for its selection and 
// creates the actual function used for the approximation
func (self *ApproxGUI) interpretFunction() error {
	switch self.function_box.SelectedValue() {
	case "sin x":
		self.function = func(x float64) float64 {
			return math.Sin(x)
		}
	case "cos x":
		self.function = func(x float64) float64 {
			return math.Cos(x)
		}
	case "e^x":
		self.function = func(x float64) float64 {
			return math.Pow(math.E, x)
		}
	case "1/sqrt(2*pi) * e^(-x^2 / 2":
		self.function = func(x float64) float64 {
			return 1.0 / math.Sqrt(2*math.Pi) * math.Pow(math.E, -x*x/2)
		}
	case "custom: ":
		// TODO
		f, err := expr.ParseExpression(self.custom_func.Text())
		if err != nil {
			self.err = err
			return err
		}
		self.function = f.Function()
	default:
		panic(fmt.Sprintf("unknown function selected: %v", self.function_box.SelectedValue()))
	}
	return nil
}

func (self *ApproxGUI) interpretAccuracy() error {
	a, err := strconv.ParseFloat(self.accuracy_box.Text(), 64)
	if err != nil {
		self.err = errors.New("Error parsing accuracy")
		return self.err
	} else {
		self.accuracy = a
		return nil
	}
}