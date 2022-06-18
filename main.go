package main

import (
	"Triton-dataseed-helper/example"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
	}))

	//router.GET("/Deletesdata", example.Deletesdata)
	router.GET("/CreateSdataFile", example.Deletesdata, example.HeadtoAppend, example.AppendGetIntPointer, example.RiskCategoryCreate, example.PolicyRuleCreate, example.PolicyRuleCategoryCreate, example.PolicyFrameworkDataCreate, example.PolicyRuleAvailableCloudDataCreate, example.AvailableCloudDataCreate, example.GuidedRemediationTemplateDataCreate, example.CloudAssetCategoryDataCreate, example.CloudAssetDataCreate, example.AppendRemediationTypeData)
	//router.GET("/AppendGetIntPointer", example.AppendGetIntPointer)
	//router.GET("/RiskCategoryCreate", example.RiskCategoryCreate)
	//router.GET("/PolicyRuleCreate", example.PolicyRuleCreate)
	//router.GET("/PolicyRuleCategoryCreate", example.PolicyRuleCategoryCreate)
	//router.GET("/PolicyFrameworkDataCreate", example.PolicyFrameworkDataCreate)
	//router.GET("/PolicyRuleAvailableCloudDataCreate", example.PolicyRuleAvailableCloudDataCreate)
	//router.GET("/AvailableCloudDataCreate", example.AvailableCloudDataCreate)
	//router.GET("/GuidedRemediationTemplateDataCreate", example.GuidedRemediationTemplateDataCreate)
	//router.GET("/CloudAssetCategoryDataCreate", example.CloudAssetCategoryDataCreate)
	//router.GET("/CloudAssetDataCreate", example.CloudAssetDataCreate)
	//router.GET("/AppendRemediationTypeData", example.AppendRemediationTypeData)
	router.Run(":5000")
}
