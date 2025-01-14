package collector

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var Collector *GradleCollector

type GradleCollector struct {
	Items []*GradleItem
}

type GradleItem struct {
	Version     string
	ReleaseTime string
	FileName    string
	FileType    string
	Sha256      string
	Sha256Url   string
	DownloadUrl string
}

func Init() {
	Collector = &GradleCollector{
		Items: getGradleAllInfo(),
	}
}

var Collector_Release_Checksums string = "https://gradle.org/release-checksums"
var Collector_Archive_Url string = "https://gradle.org/releases/"

// https://mirror.ghproxy.com/https://github.com/gradle/gradle-distributions/releases/download/v8.9.0/gradle-8.9-bin.zip
// "https://services.gradle.org/distributions/gradle-8.9-bin.zip"
// "https://services.gradle.org/distributions/gradle-8.9-bin.zip.sha256"
// https://downloads.gradle.org/distributions/gradle-8.9-bin.zip
// https://downloads.gradle.org/distributions/gradle-8.9-bin.zip.sha256
func build_GradleItem(version, version_time, sha256 string) *GradleItem {

	var item = &GradleItem{
		Version:     version,
		ReleaseTime: version_time,
		FileName:    "gradle-" + version + "-bin.zip",
		FileType:    "zip",
		Sha256:      sha256,
		//Sha256Url:   "https://downloads.gradle-dn.com/distributions/gradle-" + version + "-bin.zip.sha256",
		// DownloadUrl: "https://downloads.gradle-dn.com/distributions/gradle-" + version + "-bin.zip",
		//Sha256Url:   "https://services.gradle.org/distributions/gradle-" + version + "-bin.zip.sha256",
		//DownloadUrl: "https://services.gradle.org/distributions/gradle-" + version + "-bin.zip",
		Sha256Url:   "https://downloads.gradle.org/distributions/gradle-" + version + "-bin.zip.sha256",
		DownloadUrl: "https://downloads.gradle.org/distributions/gradle-" + version + "-bin.zip",
	}
	return item
}

func getFileNameByDownLoadUrl(url string) string {
	downloads := strings.Split(url, "/")
	file_name := downloads[len(downloads)-1]
	return file_name
}
func GetFileNameNoSuffix(file_name string) string {
	return strings.ReplaceAll(file_name, "."+getFileTypeByFileName(file_name), "")
}

func GetSha256ByUrl(url string, isGetSha256 bool) string {
	if isGetSha256 {
		resp, _ := http.Get(url)
		defer resp.Body.Close()
		bytes, _ := ioutil.ReadAll(resp.Body)
		return string(bytes)
	} else {
		return url
	}
}

func getFileTypeByFileName(filename string) string {
	filenames := strings.Split(filename, ".")
	switch filenames[len(filenames)-1] {
	case "zip":
		return "zip"
	case "gz":
		return "tar.gz"
	default:
		return filenames[len(filenames)-1]
	}
}
