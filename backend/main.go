package main

import (
	"crud_bot/db/services"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/songs/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			name := urlQuery.Get("name")
			genreName := urlQuery.Get("genre_name")
			interpreterName := urlQuery.Get("interpreter_name")
			genreId := urlQuery.Get("genre_id")
			interpreterId := urlQuery.Get("interpreter_id")

			if id != "" {
				result = services.RetrieveSongById(id)

			} else if name != "" {
				result = services.RetrieveSongByName(name)

			} else if genreName != "" {
				result = services.RetrieveSongsByGenreName(genreName)

			} else if interpreterName != "" {
				result = services.RetrieveSongByInterpreterName(interpreterName)

			} else if genreId != "" {
				result = services.RetrieveSongsByGenreId(genreId)

			} else if interpreterId != "" {
				result = services.RetrieveSongByInterpreterId(interpreterId)

			} else {
				result = services.RetrieveAllSongs()
			}

			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/genres", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			name := urlQuery.Get("name")

			if id != "" {
				result = services.RetrieveGenreById(id)

			} else if name != "" {
				result = services.RetrieveGenreByName(name)

			} else {
				result = services.RetrieveAllGenres()
			}

			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/interpreters", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			name := urlQuery.Get("name")

			if id != "" {
				result = services.RetrieveInterpreterById(id)

			} else if name != "" {
				result = services.RetrieveInterpreterByName(name)

			} else {
				result = services.RetrieveAllInterpreters()
			}

			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/topics", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			topic := urlQuery.Get("topic")

			if id != "" {
				result = services.RetrieveTopicById(id)

			} else if topic != "" {
				result = services.RetrieveTopicByTopic(topic)

			} else {
				result = services.RetrieveAllTopics()
			}

			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/subtopics", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			subtopic := urlQuery.Get("subtopic")
			topicTopic := urlQuery.Get("topic_topic")
			topicId := urlQuery.Get("topic_id")

			if id != "" {
				result = services.RetrieveSubtopicById(id)

			} else if subtopic != "" {
				result = services.RetrieveSubtopicBySubtopic(subtopic)

			} else if topicTopic != "" {
				result = services.RetrieveSubtopicByTopicTopic(topicTopic)

			} else if topicId != "" {
				result = services.RetrieveSubtopicByTopicId(topicId)

			} else {
				result = services.RetrieveAllSubtopics()
			}

			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/questions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if r.Method == "GET" {
				var result []byte

				urlQuery := r.URL.Query()
				id := urlQuery.Get("id")
				subtopicId := urlQuery.Get("subtopic_id")
				subtopicSubtopic := urlQuery.Get("subtopic_subtopic")
				topicTopic := urlQuery.Get("topic_topic")
				topicId := urlQuery.Get("topic_id")

				if id != "" {
					result = services.RetrieveQuestionById(id)

				} else if subtopicId != "" {
					result = services.RetrieveQuestionBySubtopicId(subtopicId)

				} else if subtopicSubtopic != "" {
					result = services.RetrieveQuestionBySubtopicSubtopic(subtopicSubtopic)

				} else if topicId != "" {
					result = services.RetrieveQuestionByTopicId(topicId)

				} else if topicTopic != "" {
					result = services.RetrieveQuestionByTopicTopic(topicTopic)

				} else {
					result = services.RetrieveAllQuestions()
				}

				fmt.Fprintf(w, string(result))
			}
		}
	})

	http.HandleFunc("/answers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			questionId := urlQuery.Get("question_id")

			if id != "" {
				result = services.RetrieveAnswerById(id)

			} else if questionId != "" {
				result = services.RetrieveAnswerByQuestionId(questionId)

			} else {
				result = services.RetrieveAllAnswers()
			}

			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/scores", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			botuserId := urlQuery.Get("botuser_id")

			if id != "" {
				result = services.RetrieveScoreById(id)

			} else if botuserId != "" {
				result = services.RetrieveScoreByBotuserId(botuserId)

			} else {
				result = services.RetrieveAllScores()
			}

			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var result []byte

			urlQuery := r.URL.Query()
			id := urlQuery.Get("id")
			name := urlQuery.Get("name")
			botuserId := urlQuery.Get("botuser_id")

			if id != "" {
				result = services.RetrieveItemById(id)

			} else if name != "" {
				result = services.RetrieveItemByName(name)

			} else if botuserId != "" {
				result = services.RetrieveItemByBotuserId(botuserId)

			} else {
				result = services.RetrieveAllItems()
			}

			fmt.Fprintf(w, string(result))
		}
	})

	http.HandleFunc("/botusers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := services.RetrieveAllBotusers()
			fmt.Fprintf(w, string(result))
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
