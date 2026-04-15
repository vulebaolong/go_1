package main

var roadMap = map[string]any{
	"Core": map[string]any{
		"Buổi 1 + 2 + 3": []any{
			"Biến, Hàm, Package, Struct, Mảng, Slice",
			"Erors, Defer, Groutines, Struct, Mảng, Slice",
			"Xây dựng ứng dụng quản lý chi tiêu",
		},
	},
	"CSDL": map[string]any{
		"Buổi 4 + 5": []any{
			"Mysql Với docker",
			"Relation",
		},
	},
	"Gin Framework": map[string]any{
		"Buổi 6":  "REST API & gin",
		"Buổi 7":  "Áp dụng mô hình Clean Artitecture",
		"Buổi 8":  "Interface, Eviroment & ORM Ent",
		"Buổi 9":  "Schema, Pagition & cache trong golang",
		"Buổi 10": "Filter, CORS",
	},
	"Auth": map[string]any{
		"Buổi 11": "Login, Register",
		"Buổi 12": "JWT Authentication, Middleware Protect, AccessToken",
		"Buổi 13": "RefreshToken, Login Google Authention",
	},
	"Functional": map[string]any{
		"Buổi 14": "Image Upload Local, Cloud (cloudinary)",
		"Buổi 15": "Quên mật khẩu, send mail, swagger",
		"Buổi 16": "Realtime, Chat",
		"Buổi 17": "Cache redis, Elastic search",
		"Buổi 18": "Microservice",
	},
	"Deploy": map[string]any{
		"Buổi 19": "Docker",
		"Buổi 20": "CI/CD",
	},
}
