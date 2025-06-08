package main

import (
  "encoding/json"
  "log"
  "net/http"

  "github.com/rs/cors"
  "sandhi_transformer/generic"
  "sandhi_transformer/sandhi"
  "sandhi_transformer/sandhi/core"
)

type TransformRequest struct {
  Text string `json:"text"`
}

type TransformResponse struct {
  Transformed string       `json:"transformed"`
  Chunks      []core.Chunk `json:"transformations"`
}

func transformHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w, "Only POST supported", http.StatusMethodNotAllowed)
    return
  }

  var req TransformRequest
  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    http.Error(w, "Invalid JSON", http.StatusBadRequest)
    return
  }

  normalized := generic.GenericTransforms(req.Text)
  // log.Println(normalized)
  chunks := sandhi.ApplyAllSandhiChunks(normalized)
  transformed := ""
  for _, ch := range chunks {
    transformed += ch.Text
  }

  resp := TransformResponse{
    Transformed: transformed, // or the final transformed string
    Chunks:      chunks,
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(resp)
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/transform", transformHandler)

  handler := cors.AllowAll().Handler(mux)
  //cors.New(cors.Options{
  //  AllowedOrigins: []string{"https://sandhitransformer.xyz"},
  //  AllowedMethods: []string{"POST", "OPTIONS"},
  //})


  log.Println("Server running on http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", handler))
}
