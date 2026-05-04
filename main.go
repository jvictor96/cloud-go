func main() {
	// 1. Setup inicial (Equivalente ao cloudrc e tput)
	engine := Engine{
		Columns:  80, // Detectar via term.GetSize ou exec "tput cols"
		Spacing:  2,
		MaxLines: 100,
	}

	// 2. Carregar dimensões e artes
	// Aqui você abriria o arquivo ~/.cloud/dimensions e daria Parse nas structs

	// 3. Captura o Buffer (O seu $($1) no Shell)
	// cmd := exec.Command(os.Args[1]) ... captura o output ...

	// 4. Lógica de posicionamento
	for engine.PlaceImages() {}

	// 5. Loop de Animação
	fmt.Print("\033[s") // Save cursor position
	
	maxFrames := 20 // Calcule baseado nas alturas das artes
	for f := 0; f <= maxFrames; f++ {
		engine.ManipulateBuffer(f)
		
		fmt.Print("\033[u") // Restore cursor position
		for _, line := range engine.FinalBuffer {
			fmt.Println(line)
		}
		
		time.Sleep(200 * time.Millisecond)
		// Move cursor de volta para o topo da animação
		fmt.Printf("\033[%dA", len(engine.FinalBuffer))
	}
}