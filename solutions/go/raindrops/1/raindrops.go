package raindrops

import "strconv"

func Convert(number int) string {
	var myString string = ""
    var i int = 0
    if (number % 3 == 0){
        myString+="Pling"
    	i++
    }
    
    if (number % 5 == 0){ 
    	myString+="Plang"
    	i++
    }
    
    if (number % 7 == 0){ 
    	myString+="Plong"
    	i++
    }
    
    if i == 0 {
		return "" + strconv.Itoa(number)
	}

	return myString
    
}
