package server

import (
	"context"
	"fmt"
	"log"

	// "log"
	"net/http"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	// "github.com/cmwylie19/media-controller/internal/log"
	// "github.com/djwackey/gitea/log"
)

var wg sync.WaitGroup

// Server wraps the Kubernetes configuration
type Server struct {
	// Logger TODO Structured logs
	Config *kubernetes.Clientset
	Label  string
	Port   string
}

func (s *Server) Init() {
	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	s.Config = config
}

func GetConfig() (*kubernetes.Clientset, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		// panic(err.Error())
		fmt.Println("Error getting config")
		return &kubernetes.Clientset{}, err
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		// panic(err.Error())
		fmt.Println("Error getting clientset")
		return &kubernetes.Clientset{}, err
	}

	return clientset, nil

}

func (s *Server) Serve(tlsKey, tlsCert, port, label string) error {

	// Update Server struct
	s.Port = port
	s.Label = label

	http.HandleFunc("/health", GetHealth)
	http.HandleFunc("/secrets", s.GetSecrets)

	if tlsKey != "" && tlsCert != "" {
		err := http.ListenAndServeTLS(":"+port, tlsCert, tlsKey, nil)
		if err != nil {

			log.Fatal("ListenAndServe: ", err)
			return err
		}
	} else {
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)

			return err
		}
	}
	return nil
}

func getSecrets(secrets_chan chan string, clientset *kubernetes.Clientset, ns string, label string) {
	str_secrets := ""
	defer wg.Done()

	secrets, err := clientset.CoreV1().Secrets(ns).List(context.TODO(), metav1.ListOptions{
		LabelSelector: label,
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d secrets in the cluster", len(secrets.Items))
	for _, secret := range secrets.Items {
		str_secrets += secret.Name + "\n"
		fmt.Printf("Secret: %s\n", secret.Name)
	}
	secrets_chan <- str_secrets
}

func (s *Server) GetSecrets(w http.ResponseWriter, req *http.Request) {
	ns := req.URL.Query().Get("namespace")

	fmt.Println("Namespace: ", ns)

	secrets := make(chan string)
	wg.Add(1)

	go getSecrets(secrets, s.Config, ns, s.Label)

	cluster_secrets := <-secrets
	wg.Wait()
	w.Write([]byte(cluster_secrets))
}

func GetHealth(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
}
