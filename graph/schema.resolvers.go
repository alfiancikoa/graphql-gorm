package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/alfiancikoa/graphql-gorm/graph/generated"
	"github.com/alfiancikoa/graphql-gorm/graph/model"
)

func (r *mutationResolver) CreateMovie(ctx context.Context, input model.InputMovie) (*model.Movie, error) {
	movie := model.Movie{
		Title: input.Title,
		Year:  input.Year,
	}
	if err := r.DB.Create(&movie).Error; err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}
	for _, star := range input.Stars {
		stars := model.Star{MovieID: movie.ID, Name: star.Name}
		if err := r.DB.Create(&stars).Error; err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("internal server error")
		}
	}
	return &movie, nil
}

func (r *mutationResolver) UpdateMovie(ctx context.Context, movieID int, input *model.InputMovie) (*model.Movie, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMovie(ctx context.Context, movieID int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	movies := []*model.Movie{}
	tx := r.DB.Table("movies").Select(
		"movies.id, movies.title, movies.year").Find(&movies)

	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, fmt.Errorf("internal server error")
	}
	for _, movie := range movies {
		movieId := movie.ID
		stars := []*model.Star{}
		if err := r.DB.Where("movie_id = ?", movieId).Find(&stars).Error; err != nil {
			fmt.Println(tx.Error)
			return nil, fmt.Errorf("internal server error")
		}
		movie.Stars = stars
	}
	return movies, nil
}

func (r *queryResolver) Movie(ctx context.Context, id int) (*model.Movie, error) {
	movie := model.Movie{}
	if err := r.DB.Find(&movie, id).Error; err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}
	var stars []*model.Star
	if err := r.DB.Where("movie_id=?", id).Find(&stars).Error; err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}
	movie.Stars = stars
	return &movie, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func mapStarInput(starsInput []*model.InputStar, movieId int) []model.Star {
	var stars []model.Star
	for _, star := range starsInput {
		stars = append(stars, model.Star{
			MovieID: movieId,
			Name:    star.Name,
		})
	}
	return stars
}
