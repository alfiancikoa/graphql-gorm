package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/alfiancikoa/graphql-gorm/graph/generated"
	"github.com/alfiancikoa/graphql-gorm/graph/model"
)

// Fungsi untuk menambahkan movie baru
func (r *mutationResolver) CreateMovie(ctx context.Context, input model.InputMovie) (*model.Movie, error) {
	// ambil data kemudian tampung ke variable movie
	movie := model.Movie{
		Title: input.Title,
		Year:  input.Year,
	}
	// query untuk menambahkan movie baru ke dalam tabel movie pada database
	if err := r.DB.Create(&movie).Error; err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}
	// mengisi data list stars
	for _, star := range input.Stars {
		stars := model.Star{MovieID: movie.ID, Name: star.Name}
		// query untuk memasukkan list stars pada movie tersebut ke dalam tabel star pada database
		if err := r.DB.Create(&stars).Error; err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("internal server error")
		}
	}
	// kembalikan data movie yang telah dimasukkan pada database
	return &movie, nil
}

func (r *mutationResolver) UpdateMovie(ctx context.Context, movieID int, input *model.InputMovie) (*model.Movie, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMovie(ctx context.Context, movieID int) (bool, error) {
	if err := r.DB.Delete(&model.Movie{}, movieID).Error; err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("internal server error")
	}
	if err := r.DB.Where("movie_id=?", movieID).Delete(&model.Star{}).Error; err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("internal server error")
	}
	return true, nil
}

// Fungsi untuk menampilkan seluruh list movie dan stars berdasarkan movie
func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	// variabel untuk menampung data movie
	movies := []*model.Movie{}
	// query untuk mengambil semua data movie pada tabel movie di database
	tx := r.DB.Table("movies").Select(
		"movies.id, movies.title, movies.year").Find(&movies)

	// jika ada error, maka kembalikan status error
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, fmt.Errorf("internal server error")
	}
	// proses pengambilan data stars berdasarkan movie masing-masing pada database
	for _, movie := range movies {
		// variabel untuk menampung list stars
		stars := []*model.Star{}
		// query untuk mencari data stars berdasarkan movie-nya
		if err := r.DB.Where("movie_id = ?", movie.ID).Find(&stars).Error; err != nil {
			fmt.Println(tx.Error)
			return nil, fmt.Errorf("internal server error")
		}
		// kembalikan data stars sesuai dengan movie-nya
		movie.Stars = stars
	}
	// kembalikan data seluruh list movie beserta data stars-nya
	return movies, nil
}

// Fungsi untuk menampilkan atau Get data movie berdasarkan id movie-nya
func (r *queryResolver) Movie(ctx context.Context, id int) (*model.Movie, error) {
	// variabel untuk menampung data movie
	movie := model.Movie{}
	// query database pada tabel movie untuk mencari movie by id kemudian data disimpan ke variable movie
	if err := r.DB.Find(&movie, id).Error; err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}
	// variabel untuk menampung data stars
	var stars []*model.Star
	// query database pada tabel stars untuk mencari list stars sesuai dengan id movie yang dicari
	// kemudian data disimpan ke variable stars
	if err := r.DB.Where("movie_id=?", id).Find(&stars).Error; err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}
	// assign data movie stars dengan data list stars
	movie.Stars = stars
	// kembalikan data movie by id
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
