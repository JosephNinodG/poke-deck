package domain

type GetUserCollectionRequest struct {
	UserID       int
	CollectionID int
}

type CreateUserCollectionRequest struct {
	UserID         int
	CollectionName string
}

type DbCard struct {
	ID   int
	Card PokemonCard
}
