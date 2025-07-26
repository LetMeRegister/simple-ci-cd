package add

import "testing"

func TestAdd(t *testing.T) {
	type testCase struct {
		first int 
		second int 
		result int 
	}

	cases := []testCase{
		{0,0,0,},
		{1,1,2,},
		{1,-1,0,},
		{-1,1,0,},
		{-1,-1,-2,},
	}

	for _ , testCase := range cases {
		t.Run("case" , func (t *testing.T )  {
			result := Add( testCase.first, testCase.second)
			if result != testCase.result{
				t.FailNow()
			}
		})
	}

}
