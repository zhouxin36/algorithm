package main

func main() {

}
func destCity(paths [][]string) string {
	pm := map[string]bool{}
	for _, v := range paths {
		pm[v[0]] = true
	}
	for _, v := range paths {
		if !pm[v[1]] {
			return v[1]
		}
	}
	return ""
}
