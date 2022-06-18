package example

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

type RiskCategory struct {
	Name string
}

type PolicyRule struct {
	policyFrameworkId              string
	remediationTypeId              string
	guidedRemediationTemplateId    string
	automatedRemediationTemplateId string
	policyRuleCategoryId           string
	policyNo                       string
	definiton                      string
	riskCategoryId                 string
	severity                       string
	tritonConfigRuleRef            string
	scheduleType                   string
	isEnabled                      string
}

type PolicyRuleCategory struct {
	name        string
	description string
}

type PolicyFramework struct {
	code         string
	name         string
	description  string
	isenabled    string
	iscustomized string
	logo         string
}

type PolicyRuleAvailableCloud struct {
	policyruleId string
	cloudAssetId string
}

type AvailableCloud struct {
	cloudname string
	imageurl  string
}

type RemediationType struct {
	name string
}

type GuidedRemediationTemplate struct {
	step string
}

type AutomatedRemediationTemplate struct {
	templatekey  string
	templatename string
}

type CloudAsset struct {
	name                 string
	description          string
	availablecloudId     string
	cloudAssetcategoryId string
}

type CloudAssetCategory struct {
	name string
}

func AppendRemediationTypeData(c *gin.Context) {

	//write to file
	//file, err := os.Create("CloudAsset.go")
	file, err := os.OpenFile("seedfile/sdata.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		file.WriteString("\n" + "var RemediationTypeData = []*dao.RemediationType{" + "\n")
		file.WriteString("\t" + "{" + "\n" + "\t" + "Name: remediation.Automated.String()," + "\n" + "\t" + "}," + "\n" + "\t" + "{" + "\n" + "\t" + "Name: remediation.Guided.String()," + "\n" + "\t" + "}," + "\n" + "}" + "\n")
		//fmt.Println()

	}
	//file.WriteString(string(RiskcategoryJson))
	//fmt.Println("Done")

	file.Close()
	c.JSON(200, gin.H{"message": "success"})
	//fmt.Println(string(RiskcategoryJson))

}

func AppendGetIntPointer(c *gin.Context) {

	//write to file
	//file, err := os.Create("CloudAsset.go")
	file, err := os.OpenFile("seedfile/sdata.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		file.WriteString("\n" + "func GetIntPointer(value int) *int {" + "\n")
		file.WriteString("\t" + "return &value" + "\n" + "}" + "\n")
		//fmt.Println()
	}
	//file.WriteString(string(RiskcategoryJson))
	//fmt.Println("Done")

	file.Close()
	c.JSON(200, gin.H{"message": "success"})
	//fmt.Println(string(RiskcategoryJson))

}

func Deletesdata(c *gin.Context) {
	// Removing file from the directory
	// Using Remove() function
	e := os.Remove("seedfile/sdata.go")
	if e != nil {
		//log.Fatal(e)
		myfile, e := os.Create("seedfile/sdata.go")
		if e != nil {
			log.Fatal(e)
			c.JSON(500, gin.H{"error": e.Error()})
			return
		}
		log.Println(myfile)
		myfile.Close()
		c.JSON(200, gin.H{"message": "success"})
	}

}

/*
func HeadtoAppend(c *gin.Context) {

	//write to file
	//file, err := os.Create("CloudAsset.go")
	file, err := os.OpenFile("seedfile/sdata.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		file.WriteString("package sdata" + "\n")
		file.WriteString("\n" + "\n" + "import (" + "\n" + "\t" + "\"" + "github.com/DygiSecure/triton-microservice-backend/services/policy/internal/pkg/dao" + "\"" + "\n" + "\t" + "\"" + "github.com/DygiSecure/triton-microservice-backend/services/policy/pkg/remediation" + "\"" + "\n" + "\t" + "\"" + "github.com/DygiSecure/triton-microservice-backend/services/policy/pkg/severity" + "\"" + "\n" + ")" + "\n")
		//fmt.Println()
	}
	//file.WriteString(string(RiskcategoryJson))
	//fmt.Println("Done")

	file.Close()
	//fmt.Println(string(RiskcategoryJson))
	c.JSON(200, gin.H{"message": "success"})
}
*/

func CloudAssetDataCreate(c *gin.Context) {
	csvFile, errc := os.Open("C:/Users/USER/OneDrive/Desktop/Triton-dataseed-helper/Policy Compliances  - Sample Data Feed v1.0.0.xlsx - cloudAsset.csv")
	if errc != nil {
		log.Fatalf("failed opening file: %s", errc)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Read column label data and discard
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	var cloudassetData []CloudAsset
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		cloudassetData = append(cloudassetData, CloudAsset{
			availablecloudId:     line[1],
			cloudAssetcategoryId: line[2],
			name:                 line[3],
			description:          line[4],
		})
	}

	RiskcategoryJson, err := json.MarshalIndent(cloudassetData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	bytes := []byte(RiskcategoryJson)
	json.Unmarshal(bytes, &cloudassetData)

	file, err := os.OpenFile("seedfile/sdata.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		file.WriteString("\n" + "var CloudAssetData = []*dao.CloudAsset{" + "\n")
		for k := range cloudassetData {
			AvailableCloudID := ("{\n" + "\t" + "AvailableCloudID" + ":" + string(cloudassetData[k].availablecloudId) + "," + "\n")
			CloudAssetCategoryID := ("\t" + "CloudAssetCategoryID" + ":" + string(cloudassetData[k].cloudAssetcategoryId) + "," + "\n")
			Name := ("\t" + "Name" + ":" + "\"" + string(cloudassetData[k].name) + "\"" + "," + "\n")
			Description := ("\t" + "Description" + ":" + "\"" + string(cloudassetData[k].description) + "\"" + "," + "\n},")

			//fmt.Printf("{"+"\t"+"AvailableCloudID : %q"+","+"\n"+"\t"+"CloudAssetCategoryID: %q"+","+"\n"+"\t"+"Name: %q"+","+"\n"+"\t"+"Description: %q"+","+"\n},", string(cloudassetData[k].availablecloudId), string(cloudassetData[k].cloudAssetcategoryId), string(cloudassetData[k].name), string(cloudassetData[k].description))
			file.WriteString(AvailableCloudID + CloudAssetCategoryID + Name + Description + "\n")

		}
		file.WriteString("}" + "\n")

	}
	file.Close()

	c.JSON(200, gin.H{"message": "success"})
}

func CloudAssetCategoryDataCreate(c *gin.Context) {
	csvFile, errc := os.Open("C:/Users/USER/OneDrive/Desktop/Triton-dataseed-helper/Policy Compliances  - Sample Data Feed v1.0.0.xlsx - cloudAssetCategory.csv")

	if errc != nil {
		log.Fatalf("failed opening file: %s", errc)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Read column label data and discard
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	var cloudassetcategorydata []CloudAssetCategory
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		cloudassetcategorydata = append(cloudassetcategorydata, CloudAssetCategory{
			name: line[1],
		})
	}

	RiskcategoryJson, err := json.MarshalIndent(cloudassetcategorydata, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	bytes := []byte(RiskcategoryJson)
	json.Unmarshal(bytes, &cloudassetcategorydata)

	//write to file
	//file, err := os.Create("CloudAssetCategory.go")
	file, err := os.OpenFile("seedfile/sdata.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {

		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		file.WriteString("\n" + "var CloudAssetCategoryData = []*dao.CloudAssetCategory{" + "\n")
		for k := range cloudassetcategorydata {
			Step := ("{\n" + "\t" + "Name" + ":" + "\"" + string(cloudassetcategorydata[k].name) + "\"" + "," + "\n},")

			//fmt.Printf("{"+"\t"+"Name : %q"+","+"\n},", string(cloudassetcategorydata[k].name))
			file.WriteString(Step + "\n")
			//fmt.Println()
		}
		file.WriteString("}" + "\n")

	}
	file.Close()

	c.JSON(200, gin.H{"message": "success"})
}

func GuidedRemediationTemplateDataCreate(c *gin.Context) {
	csvFile, errc := os.Open("C:/Users/USER/OneDrive/Desktop/Triton-dataseed-helper/Policy Compliances  - Sample Data Feed v1.0.0.xlsx - GuidedRemediationTemplate.csv")

	if errc != nil {
		log.Fatalf("failed opening file: %s", errc)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Read column label data and discard
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	var guidedremediationtemplateData []GuidedRemediationTemplate
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		guidedremediationtemplateData = append(guidedremediationtemplateData, GuidedRemediationTemplate{
			step: line[1],
		})
	}

	RiskcategoryJson, err := json.MarshalIndent(guidedremediationtemplateData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	bytes := []byte(RiskcategoryJson)
	json.Unmarshal(bytes, &guidedremediationtemplateData)

	file, err := os.OpenFile("seedfile/sdata.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		file.WriteString("\n" + "var GuidedRemediationTemplateData = []*dao.GuidedRemediationTemplate{" + "\n")
		for k := range guidedremediationtemplateData {
			Step := ("{\n" + "\t" + "Step" + ":" + string(guidedremediationtemplateData[k].step) + "\n},")

			//fmt.Printf("{"+"\t"+"Step : %q"+","+"\n},", string(guidedremediationtemplateData[k].step))
			file.WriteString(Step + "\n")
			//fmt.Println()
		}
		file.WriteString("}" + "\n")

	}
	file.Close()

	c.JSON(200, gin.H{"message": "success"})
}

func AvailableCloudDataCreate(c *gin.Context) {
	csvFile, errc := os.Open("C:/Users/USER/OneDrive/Desktop/Triton-dataseed-helper/Policy Compliances  - Sample Data Feed v1.0.0.xlsx - AvailableCloud.csv")
	if errc != nil {
		log.Fatalf("failed opening file: %s", errc)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Read column label data and discard
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	var availablecloudData []AvailableCloud
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		availablecloudData = append(availablecloudData, AvailableCloud{
			cloudname: line[1],
			imageurl:  line[2],
		})
	}

	RiskcategoryJson, err := json.MarshalIndent(availablecloudData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	bytes := []byte(RiskcategoryJson)
	json.Unmarshal(bytes, &availablecloudData)

	file, err := os.OpenFile("seedfile/sdata.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		file.WriteString("\n" + "var AvailableCloudData = []*dao.AvailableCloud{" + "\n")
		for k := range availablecloudData {
			CloudName := ("{\n" + "\t" + "CloudName" + ":" + "\"" + string(availablecloudData[k].cloudname) + "\"" + "," + "\n")
			ImageURL := ("\t" + "ImageURL" + ":" + "\"" + string(availablecloudData[k].imageurl) + "\"" + "," + "\n},")

			//fmt.Printf("{"+"\t"+"CloudName : %q"+","+"\n"+"\t"+"ImageURL: %q"+","+"\n},", string(availablecloudData[k].cloudname), string(availablecloudData[k].imageurl))
			file.WriteString(CloudName + ImageURL + "\n")
			//fmt.Println()
		}
		file.WriteString("}" + "\n")

	}
	file.Close()
	//fmt.Println(string(RiskcategoryJson))
	c.JSON(200, gin.H{"message": "success"})
}

func PolicyRuleAvailableCloudDataCreate(c *gin.Context) {
	csvFile, errc := os.Open("C:/Users/USER/OneDrive/Desktop/Triton-dataseed-helper/Policy Compliances  - Sample Data Feed v1.0.0.xlsx - PolicyRuleAvailableCloud.csv")

	if errc != nil {
		log.Fatalf("failed opening file: %s", errc)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Read column label data and discard
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	var policyruleavailableData []PolicyRuleAvailableCloud
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		policyruleavailableData = append(policyruleavailableData, PolicyRuleAvailableCloud{
			policyruleId: line[1],
			cloudAssetId: line[2],
		})
	}

	RiskcategoryJson, err := json.MarshalIndent(policyruleavailableData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	bytes := []byte(RiskcategoryJson)
	json.Unmarshal(bytes, &policyruleavailableData)

	file, err := os.OpenFile("seedfile/sdata.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		file.WriteString("\n" + "var PolicyRuleAvailableCloudData = []*dao.PolicyRuleAvailableCloud{" + "\n")
		for k := range policyruleavailableData {
			PolicyRuleID := ("{\n" + "\t" + "PolicyRuleID" + ":" + string(policyruleavailableData[k].policyruleId) + "," + "\n")
			CloudAssetID := ("\t" + "CloudAssetID" + ":" + string(policyruleavailableData[k].cloudAssetId) + "," + "\n},")

			//fmt.Printf("{"+"\t"+"PolicyRuleID : %q"+","+"\n"+"\t"+"AvailableCloudID: %q"+","+"\n},", string(policyruleavailableData[k].policyruleId), string(policyruleavailableData[k].availablecloudId))
			file.WriteString(PolicyRuleID + CloudAssetID + "\n")
			//fmt.Println()
		}
		file.WriteString("}" + "\n")

	}
	file.Close()

	c.JSON(200, gin.H{"message": "success"})
}

func PolicyFrameworkDataCreate(c *gin.Context) {
	csvFile, errc := os.Open("C:/Users/USER/OneDrive/Desktop/Triton-dataseed-helper/Policy Compliances  - Sample Data Feed v1.0.0.xlsx - PolicyFramework.csv")

	if errc != nil {
		log.Fatalf("failed opening file: %s", errc)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Read column label data and discard
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	var policyframeworkData []PolicyFramework
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		policyframeworkData = append(policyframeworkData, PolicyFramework{
			code:         line[1],
			name:         line[2],
			description:  line[3],
			isenabled:    line[4],
			iscustomized: line[5],
			logo:         line[6],
		})
	}

	RiskcategoryJson, err := json.MarshalIndent(policyframeworkData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	bytes := []byte(RiskcategoryJson)
	json.Unmarshal(bytes, &policyframeworkData)

	file, err := os.OpenFile("seedfile/sdata.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		file.WriteString("\n" + "var PolicyFrameworkData = []*dao.PolicyFramework{" + "\n")
		for k := range policyframeworkData {
			Code := ("{\n" + "\t" + "Code" + ":" + "\"" + string(policyframeworkData[k].code) + "\"" + "," + "\n")
			Name := ("\t" + "Name" + ":" + "\"" + string(policyframeworkData[k].name) + "\"" + "," + "\n")
			Description := ("\t" + "Description" + ":" + "\"" + string(policyframeworkData[k].description) + "\"" + "," + "\n")
			IsEnabled := ("\t" + "IsEnabled" + ":" + string(policyframeworkData[k].isenabled) + "," + "\n")
			IsCustomized := ("\t" + "IsCustomized" + ":" + string(policyframeworkData[k].iscustomized) + "," + "\n")
			Logo := ("\t" + "Logo" + ":" + "\"" + string(policyframeworkData[k].logo) + "\"" + "," + "\n},")

			//fmt.Printf("{"+"\t"+"Code : %q"+","+"\n"+"\t"+"Name: %q"+","+"\n"+"\t"+"Description: %q"+","+"\n"+"\t"+"IsEnabled: %q"+","+"\n"+"\t"+"IsCustomized: %q"+","+"\n"+"\t"+"Logo: %q"+","+"\n},", string(policyframeworkData[k].code), string(policyframeworkData[k].name), string(policyframeworkData[k].description), string(policyframeworkData[k].isenabled), string(policyframeworkData[k].iscustomized), string(policyframeworkData[k].logo))
			file.WriteString(Code + Name + Description + IsEnabled + IsCustomized + Logo + "\n")
			//fmt.Println()
		}
		file.WriteString("}" + "\n")

	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	gocsv.UnmarshalBytes(fileBytes, &policyframeworkData)
	c.JSON(200, policyframeworkData)
}

func PolicyRuleCategoryCreate(c *gin.Context) {
	csvFile, errc := os.Open("C:/Users/USER/OneDrive/Desktop/Triton-dataseed-helper/Policy Compliances  - Sample Data Feed v1.0.0.xlsx - PolicyRuleCategory.csv")

	if errc != nil {
		log.Fatalf("failed opening file: %s", errc)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Read column label data and discard
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	var policycategoryData []PolicyRuleCategory
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		policycategoryData = append(policycategoryData, PolicyRuleCategory{
			name:        line[1],
			description: line[2],
		})
	}

	RiskcategoryJson, err := json.MarshalIndent(policycategoryData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	bytes := []byte(RiskcategoryJson)
	json.Unmarshal(bytes, &policycategoryData)

	file, err := os.OpenFile("seedfile/sdata.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {

		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		file.WriteString("\n" + "var PolicyRuleCategoryData = []*dao.PolicyRuleCategory{" + "\n")
		for k := range policycategoryData {
			Name := ("{\n" + "\t" + "Name" + ":" + "\"" + string(policycategoryData[k].name) + "\"" + "," + "\n")
			Description := ("\t" + "Description" + ":" + "\"" + string(policycategoryData[k].description) + "\"" + "," + "\n},")

			//fmt.Printf("{"+"\t"+"Name : %q"+","+"\n"+"\t"+"Description: %q"+","+"\n},", string(policycategoryData[k].name), string(policycategoryData[k].description))
			file.WriteString(Name + Description + "\n")
			//fmt.Println()
		}
		file.WriteString("}" + "\n")

	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	gocsv.UnmarshalBytes(fileBytes, &policycategoryData)
	c.JSON(200, policycategoryData)
}

func PolicyRuleCreate(c *gin.Context) {
	csvFile, errc := os.Open("C:/Users/USER/OneDrive/Desktop/Triton-dataseed-helper/Policy Compliances  - Sample Data Feed v1.0.0.xlsx - PolicyRule.csv")

	if errc != nil {
		log.Fatalf("failed opening file: %s", errc)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Read column label data and discard
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	var policyData []PolicyRule
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		policyData = append(policyData, PolicyRule{
			policyFrameworkId:              line[1],
			remediationTypeId:              line[2],
			guidedRemediationTemplateId:    line[3],
			automatedRemediationTemplateId: line[4],
			policyRuleCategoryId:           line[5],
			policyNo:                       line[6],
			definiton:                      line[7],
			riskCategoryId:                 line[8],
			severity:                       line[9],
			tritonConfigRuleRef:            line[10],
			scheduleType:                   line[11],
			isEnabled:                      line[12],
		})
	}

	RiskcategoryJson, err := json.MarshalIndent(policyData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	bytes := []byte(RiskcategoryJson)
	json.Unmarshal(bytes, &policyData)

	//write to file
	//file, err := os.Create("PolicyRule.go")
	file, err := os.OpenFile("seedfile/sdata.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		file.WriteString("\n" + "var PolicyRuleData = []*dao.PolicyRule{" + "\n")
		for k := range policyData {
			policyFrameworkId := ("{\n" + "\t" + "PolicyFrameworkID" + ":" + string(policyData[k].policyFrameworkId) + "," + "\n")
			remediationTypeId := ("\t" + "RemediationTypeID" + ":" + string(policyData[k].remediationTypeId) + "," + "\n")
			guidedRemediationTemplateId := ("\t" + "GuidedRemediationTemplateID" + ":" + string(policyData[k].guidedRemediationTemplateId) + "," + "\n")
			//automatedRemediationTemplateId := ("\t" + "AutomatedRemediationTemplateID" + ":" + "\"" + string(policyData[k].automatedRemediationTemplateId) + "\"" + "," + "\n")
			policyRuleCategoryId := ("\t" + "PolicyRuleCategoryID" + ":" + string(policyData[k].policyRuleCategoryId) + "," + "\n")
			policyNo := ("\t" + "PolicyNo" + ":" + "\"" + string(policyData[k].policyNo) + "\"" + "," + "\n")
			definiton := ("\t" + "Definition" + ":" + "\"" + string(policyData[k].definiton) + "\"" + "," + "\n")
			riskCategoryId := ("\t" + "RiskCategoryID" + ":" + string(policyData[k].riskCategoryId) + "," + "\n")
			severity := ("\t" + "Severity" + ":" + "severity." + string(policyData[k].severity) + ".String()" + "," + "\n")
			tritonConfigRuleRef := ("\t" + "TritonConfigRuleRef" + ":" + "\"" + string(policyData[k].tritonConfigRuleRef) + "\"" + "," + "\n")
			scheduleType := ("\t" + "ScheduleType" + ":" + "\"" + string(policyData[k].scheduleType) + "\"" + "," + "\n")
			isEnabled := ("\t" + "IsEnabled" + ":" + string(policyData[k].isEnabled) + "," + "\n},")

			//String--->int conversion not need
			//policyFrameworkId_number, err := strconv.ParseUint(string(policyData[k].policyFrameworkId), 10, 32)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			//	policyFrameworkId_IntNum := int(policyFrameworkId_number)

			//fmt.Printf("{"+"\t"+"PolicyFrameworkID : %d"+","+"\n"+"\t"+"RemediationTypeID: %q"+","+"\n"+"\t"+"GuidedRemediationTemplateId: %q"+","+"\n"+"\t"+"AutomatedRemediationTemplateId: %q"+","+"\n"+"\t"+"CloudAssetID: %q"+","+"\n"+"\t"+"PolicyRuleCategoryID: %q"+","+"\n"+"\t"+"PolicyNo: %q"+","+"\n"+"\t"+"Definition: %q"+","+"\n"+"\t"+"RiskCategoryID: %q"+","+"\n"+"\t"+"Severity: %q"+","+"\n"+"\t"+"TritonConfigRuleRef: %q"+","+"\n"+"\t"+"ScheduleType: %q"+","+"\n"+"\t"+"IsEnabled: %q"+","+"\n},", policyFrameworkId_IntNum, string(policyData[k].remediationTypeId), string(policyData[k].guidedRemediationTemplateId), string(policyData[k].automatedRemediationTemplateId), string(policyData[k].cloudAssetId), string(policyData[k].policyRuleCategoryId), string(policyData[k].policyNo), string(policyData[k].definiton), string(policyData[k].riskCategoryId), string(policyData[k].severity), string(policyData[k].tritonConfigRuleRef), string(policyData[k].scheduleType), string(policyData[k].isEnabled))
			file.WriteString(policyFrameworkId + remediationTypeId + guidedRemediationTemplateId + policyRuleCategoryId + policyNo + definiton + riskCategoryId + severity + tritonConfigRuleRef + scheduleType + isEnabled + "\n")

		}
		file.WriteString("}" + "\n")
	}

	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	gocsv.UnmarshalBytes(fileBytes, &policyData)
	c.JSON(200, policyData)
}

func RiskCategoryCreate(c *gin.Context) {
	csvFile, errc := os.Open("C:/Users/USER/OneDrive/Desktop/Triton-dataseed-helper/Policy Compliances  - Sample Data Feed v1.0.0.xlsx - RiskCategory.csv")

	if errc != nil {
		log.Fatalf("failed opening file: %s", errc)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Read column label data and discard
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	var riskData []RiskCategory
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		found := false
		for i := range riskData {
			if riskData[i].Name == line[1] {
				found = true

				break
			}
		}
		if !found {
			riskData = append(riskData, RiskCategory{
				Name: line[1],
			})
		}
	}

	RiskcategoryJson, err := json.MarshalIndent(riskData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	bytes := []byte(RiskcategoryJson)
	json.Unmarshal(bytes, &riskData)

	file, err := os.OpenFile("seedfile/sdata.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		file.WriteString("\n" + "var RiskCategoryData = []*dao.RiskCategory{" + "\n")
		for l := range riskData {
			name := ("{\n" + "\t" + "Name" + ":" + "\"" + string(riskData[l].Name) + "\"" + "," + "\n}")
			//fmt.Printf("{\n"+"\t"+"Name : %q"+"\n}", string(riskData[l].Name))
			file.WriteString(name + "," + "\n")

		}
		file.WriteString("}" + "\n")

	}

	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	gocsv.UnmarshalBytes(fileBytes, &riskData)
	c.JSON(200, riskData)
}
