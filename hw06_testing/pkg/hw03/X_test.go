package packageX

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		name string
		x    int
		y    int
	}{
		{
			name: "1",
			x:    123,
			y:    532,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			for i := 0; i < tC.y; i++ {
				if i%2 == 0 {
					for index := 0; index < tC.x; index++ {
						X(index, tC.x)

					}
				} else {
					for index := 1; index < tC.x; index++ {
						X(index, tC.x)
					}
				}
			}
		})

	}
}
