package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Project Info
type ProjectInfo struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Detail string   `json:"detail" gorm:"type:longtext"`
	Goal string `json:"goal"`
	Github_Title string `json:"github_title"`
	Github_Url string `json:"url"`
	Others_Place string `json:"others_place"`
	Period string `json:"period"`
	Img1, Img2, Img3, Img4, Img5, Img6, Img7, Img8, Img9, Img10 string
}

func Init()  {
	db, err := gorm.Open("mysql", "root:paul654321@(127.0.0.1:8889)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&ProjectInfo{})
	db.Delete(&ProjectInfo{})

	//mes project
	project1 := ProjectInfo{
		ID: 1, Name : "Manufacturing Execution System (MES)",
		Detail :  "本系統從設計到開發皆實際調查現場及客戶需求，\n" +
			"也於系統正式上線後，進行維護及協調之工作，\n"+
			"除了現場人員實際使用之外，\n" +
			"同時也有國外客戶採此系統，\n"+
			"也很榮幸於2019 TIMTOS工具機展上，\n" +
			"展出實際派單功能。\n" +
			"\n" +
			"# 以PowerBuilder開發之輕量化系統\n" +
			"# 一鍵派單功能，有效提升整體系統速度\n" +
			"# 採GUI介面設計\n" ,
		Goal : "We want to speed up processes with work order dispatch and design GUI  makes it easy to use.",
		Github_Title: "",
		Github_Url:"#",
		Others_Place: "videos",
		Img1:  "/public/img/content/744/portfolio/mes2.jpg",
		Img2:  "/public/img/content/744/portfolio/m1.jpg",
		Img3:  "/public/img/content/744/portfolio/m2.jpg",
		Img4:  "/public/img/content/744/portfolio/m3.jpg",
		Img5:  "/public/img/content/744/portfolio/m4.jpg",
		Period: "2018.12 ~ 2019.05",
	}

	//ERP-SFT project
	project2 := ProjectInfo{
		ID: 2, Name : "ERP/SFT Learning System",
		Detail :  "主要設計一套以動畫教學影片為主題之學習網站，\n" +
			"為了改善中小企業在導入ERP或SFT時，\n" +
			"能加快其導入之時間，以降低成本，\n" +
			"故採較趣味之動畫影片為題材，\n" +
			"另外以『蛋黃酥製作過程』為案例進行教學，\n" +
			"同時結合Arduino操控六軸機器手臂，\n" +
			"來模擬工廠之生產流程。\n" +
			"\n" +
			"# 動畫影片以CrazyTalk Animator3製作\n" +
			"# Web以Laravel Framework搭配MVC架構開發\n" +
			"# 榮獲資管系專題成果展 第一名\n" ,
		Goal: "We want to make users learning efficiency and provided some learning videos about animation.",
		Github_Title: "ERP-SFT-Learning-System",
		Github_Url:"https://github.com/SUN06444/ERP-SFT-Learning-System",
		Img1:  "/public/img/content/744/portfolio/Seminar.jpg",
		Img2: "/public/img/content/744/portfolio/s2.jpg",
		Img3: "/public/img/content/744/portfolio/s3.jpg",
		Img4: "/public/img/content/744/portfolio/s4.jpg",
		Img5: "/public/img/content/744/portfolio/s5.jpg",
		Period: "2018.02 ~ 2018.06",
	}

	//Renthouse web project
	project3 := ProjectInfo{
		ID: 3, Name : "Renthouse web",
		Detail : "專為大專生設計之租屋平台，\n" +
			"提供房屋刊登、收藏、比較等功能，\n" +
			"同時結合Google Map Api，\n" +
			"用戶能分區查詢附近之醫療機構、診所等區域，\n" +
			"從系統分析到整個功能構想及後台管理，\n" +
			"算第一次正式寫一個完整的網站，\n" +
			"也獲得了相關之系統實務經驗，\n" +
			"於github上也有更詳細之安裝等說明過程。\n"+
			"\n" +
			"# Web以Laravel Framework搭配MVC架構開發\n" +
			"# 榮獲資管系企業應用程式設計競賽 第一名",
		Goal: "We want to provide students a safe platform to find accommodation easily.",
		Github_Title: "RenthouseWeb-Laravel",
		Github_Url:"https://github.com/SUN06444/RenthouseWeb-Laravel",
		Img1:  "/public/img/content/744/portfolio/renthouse.jpg",
		Img2:  "/public/img/content/744/portfolio/r2.jpg",
		Img3:  "/public/img/content/744/portfolio/r3.jpg",
		Period: "2017.11 ~ 2018.01",
	}

	//IoT Remote Control (iOS App) project
	project4 := ProjectInfo{
		ID: 4, Name : "IoT Remote Control (iOS App)</h3>",
		Detail : "此App主要用來搭配實際搭建之虛擬工廠，\n" +
			"搭配ESP8266 WiFi晶片作為遠端操控之目的，\n" +
			"左圖有App整個頁面架構圖，\n" +
			"本專案重點在於如何運用iOS App去搭配硬體設備，\n" +
			"\n" +
			"# 以Xcode作為開發iOS App之環境\n"+
			"# iOS App控制ESP8266晶片",
		Github_Title: "Food_Factory_IoT_App",
		Github_Url:"https://github.com/SUN06444/Food_Factory_IoT_App",
		Img1:  "/public/img/content/744/iOS/ios.jpg",
		Img2:  "/public/img/content/744/iOS/ios1.jpg",
		Img3:  "/public/img/content/744/iOS/ios2.jpg",
		Period: "2018.04 ~ 2018.05",
	}

	//IoT Factory
	project5 := ProjectInfo{
		ID: 5, Name: "IoT Factory</h3>",
		Detail : "主要用來實際搭建虛擬工廠，\n" +
			"搭配ESP8266 WiFi晶片作為遠端操控之目的，\n" +
			"左圖有虛擬工廠設計架構圖，\n" +
			"其中的機器手臂，也因別的工管系之合作專案，\n" +
			"有進一步之認識。\n" +
			"而本專案之重點在於整體線路之搭配，\n" +
			"並同時搭配伺服馬達驅動及輸送帶運轉等項目，\n" +
			"故搭載多組繼電器，來維持其穩定度，\n" +
			"\n" +
			"# 以Arduino控制六軸機器手臂\n"+
			"# Relay + DC、AC Motor + Robot Arm\n",
		Github_Title: "",
		Github_Url:"#",
		Img1:  "/public/img/content/744/iOS/factory.jpg",
		Img2:  "/public/img/content/744/iOS/IoT1.jpg",
		Img3:  "/public/img/content/744/iOS/IoT2.jpg",
		Img4:  "/public/img/content/744/iOS/IoT3.jpg",
		Img5:  "/public/img/content/744/iOS/IoT3.jpg",
		Img6:  "/public/img/content/744/iOS/ifactory_all.jpg",
		Img7:  "/public/img/content/744/iOS/all.jpg",
		Period: "2018.02 ~ 2018.04",
	}

	//ESP8266+DHT11 (Firebase) project
	project6 := ProjectInfo{
		ID: 6, Name: "ESP8266+DHT11 (Firebase)</h3>",
		Detail: "以DHT11偵測溫濕度，\n" +
			"並將數據傳到Firebase資料庫，\n" +
			"再透過Fireboard製作數據圖，\n" +
			"於github上還有其他ESP8266搭配之項目\n" +
			"\n" +
			"# ESP8266+LED+Servo+DHT11+Firebase",
		Github_Title: "ESP8266",
		Github_Url:"https://github.com/SUN06444/ESP8266",
		Img1:  "/public/img/content/744/arduino/CircuitLayout1.jpg",
		Img2:  "/public/img/content/744/arduino/ESP8266/e1.jpg",
		Img3:  "/public/img/content/744/arduino/ESP8266/e2.jpg",
		Img4:  "/public/img/content/744/arduino/ESP8266/e3.jpg",
		Period: "2018.01 ~ 2018.02",
	}

	//Arduino+Bluetooth(HC06) project
	project7 := ProjectInfo{
		ID: 7, Name: "Arduino+Bluetooth(HC06)</h3>",
		Detail: "使用HC06(藍芽模組)與CoolTerm控制多個LED燈",
		Github_Title: "Bluetooth-CoolTerm-For-Mac-",
		Github_Url:"https://github.com/SUN06444/Bluetooth-CoolTerm-For-Mac-",
		Img1:  "/public/img/content/744/arduino/CircuitLayout2.jpg",
		Period: "2018.01 ~ 2018.02",
	}

	//PX-mart Management and App Design project
	project8 := ProjectInfo{
		ID: 8, Name: "PX-mart Management and App Design</h3>",
		Detail: "針對全聯網購做的系統分析與設計，\n" +
			"針對APP網購介面及Web後端管理系統做設計，\n" +
			"並考量使用者體驗，設計使用者之介面及內容。\n"+
			"\n"+
			"# APP使用Fluid-Ui做介面設計\n" +
			"# Web使用Bootrap+jQuery做設計",
		Github_Title: "PXmart",
		Github_Url:"https://github.com/SUN06444/PXmart",
		Img1:  "/public/img/content/744/uiux/pxmart/2.jpg",
		Img2:  "/public/img/content/744/uiux/pxmart/1.jpg",
		Img3:  "/public/img/content/744/uiux/pxmart/4.jpg",
		Img4:  "/public/img/content/744/uiux/pxmart/5.jpg",
		Img5:  "/public/img/content/744/uiux/pxmart/3.jpg",
		Img6:  "/public/img/content/744/uiux/pxmart/app1.jpg",
		Img7:  "/public/img/content/744/uiux/pxmart/app2.jpg",
		Img8:  "/public/img/content/744/uiux/pxmart/app3.jpg",
		Img9:  "/public/img/content/744/uiux/pxmart/app4.jpg",
		Img10:  "/public/img/content/744/uiux/pxmart/app5.jpg",
		Period: "2017.02 ~ 2017.06",
	}

	/*
	//Monthly Report Design
	var project9 = ProjectInfo{
		ID: 9, Name: "Monthly Report Design</h3>",
		Detail: "Monthly Report Design系統開發Deatails說明",
		Github_Title: "",
		Github_Url:"#",
		Img1:  "/public/img/content/744/uiux/dm5.jpg",
		Period: "2018.08 ~ 2018.09",
	}
	*/

	db.Create(&project1)
	db.Create(&project2)
	db.Create(&project3)
	db.Create(&project4)
	db.Create(&project5)
	db.Create(&project6)
	db.Create(&project7)
	db.Create(&project8)
}

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:paul654321@(127.0.0.1:8889)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}