package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	terms := []string{
		"algoritmos", "estrutura de dados", "programacao dinamica",
		"teoria dos grafos", "automatos finitos", "linguagens formais",
		"sistemas operacionais", "compiladores", "redes de computadores",
		"banco de dados", "inteligencia artificial", "aprendizado de maquina",
		"redes neurais", "processamento de linguagem natural", "visao computacional",

		"python", "java", "golang", "javascript", "typescript",
		"c", "c++", "rust", "swift", "kotlin",

		"engenharia de software", "arquitetura de software", "microservicos",
		"design patterns", "programacao orientada a objetos", "programacao funcional",
		"teste de software", "controle de versao", "git", "github",

		"docker", "kubernetes", "aws", "azure", "google cloud", "devops",
		"ci cd", "terraform", "ansible", "monitoramento prometheus", "observabilidade grafana",

		"analise de dados", "pandas python", "numpy python", "matplotlib",
		"scikit learn", "machine learning supervisionado", "machine learning nao supervisionado",

		"computacao quantica", "seguranca da informacao", "criptografia",
		"blockchain", "web3", "realidade aumentada", "realidade virtual",
		"internet das coisas", "iot", "automacao industrial", "robotica",
	}

	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir("data", 0755)
	}

	file, err := os.Create("data/searches.txt")
	if err != nil {
		fmt.Println("[ERROR] could not create file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Generating searches dataset...")

	for _, t := range terms {
		freq := int(50 + 500*rand.Float64()*rand.Float64())
		fmt.Fprintf(file, "%s %d\n", t, freq)
	}

	fmt.Println("Dataset generate on success: data/searches.txt")
}
