package main

import (
	"github.com/skratchdot/open-golang/open"
	// "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type toTest struct {
        Browsers []string
        Urls []string
        AtOnce bool
}

func (c *toTest) getConf() *toTest {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c toTest
	c.getConf()

	// fmt.Println(c)

	for _, browsers := range c.Browsers {
		// element is the element from someSlice for where we are
		for _, link := range c.Urls {
			if (!c.AtOnce){
				open.RunWith(link, browsers)
			} else {
				open.StartWith(link, browsers)
			}
		}
	}

}
