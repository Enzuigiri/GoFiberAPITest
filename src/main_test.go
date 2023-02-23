package main

import (
	"testing"
)

func TestSimpleCalc(t *testing.T) {
	if Calc(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}
func TestCalc(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"Test 1", args{x: 5}, 7},
		{"Test 2", args{x: 8}, 10},
		{"Test 3", args{x: -2}, 0},
		{"Test 4", args{x: -9999}, -9997},
		{"Test 5", args{x: 9999}, 10001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Calc(tt.args.x); gotResult != tt.wantResult {
				t.Errorf("Calc() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// func Test_Handler(t *testing.T) {
// 	app := fiber.New()

// 	app.Get("/test", func(c *fiber.Ctx) error {
// 		return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 			"success": true,
// 		})
// 	})

// 	resp, err := app.Test(httptest.NewRequest("GET", "/test", nil))

// 	utils.AssertEqual(t, nil, err, "app.test")
// 	utils.AssertEqual(t, 200, resp.StatusCode, "status_code")

// 	body, err := io.ReadAll(resp.Body)

// 	utils.AssertEqual(t, "success", body.success, "respond_body")
// }
