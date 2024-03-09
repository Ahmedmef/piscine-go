if str1 == "" && str2=="" {
	return -1
	
}
str := str1 +str2 
count := 0
fot _, s:= range str {
	unique :=true
	for j,c:=range str {
		if s == c && i !=j{
			unique = false
		}
		if unique {
			count ++ 
		}
	}
	return count 
}