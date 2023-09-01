package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Pokemon struct {
	Pokedex_Number int
	Name           string
	Type1          string
	Type2          string
	Classification string
	Generation     int
	Legendary      bool
}

func GetPokemonByName(name string) *Pokemon {
	db, err := sql.Open("sqlite3", "./Pokemon.db")
	if err != nil {
		log.Println("Failed to open database.")
		return nil
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT Pokedex_number, Name, Type1, Type2, Classification, Generation, Legendary FROM Pokemon WHERE Name = ?")
	if err != nil {
		log.Println("Failed to prepare select statement error:", err)
		return nil
	}

	rows, err := stmt.Query(name)
	if err != nil {
		log.Println("Failed to select any rows.")
		return nil
	}
	pokemon := Pokemon{}
	for rows.Next() {
		err := rows.Scan(&pokemon.Pokedex_Number, &pokemon.Name, &pokemon.Type1, &pokemon.Type2, &pokemon.Classification, &pokemon.Generation, &pokemon.Legendary)
		if err != nil {
			return nil
		}
	}

	return &pokemon
}
