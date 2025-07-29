package add_test

import "testing"
import "add"

func TestMultiply(t *testing.T ){
	t.Run("zero by zero", func (t *testing.T )  {
		if(0 != add.Multiply(0 , 0 )){
			t.FailNow()
		}
	})
	t.Run("fail", func (t *testing.T)  {
		t.FailNow()
	})
}
