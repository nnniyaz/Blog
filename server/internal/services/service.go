package services

import (
	"github.com/nnniyaz/blog/internal/domain/base/config"
	"github.com/nnniyaz/blog/internal/repos"
	"github.com/nnniyaz/blog/internal/services/article"
	"github.com/nnniyaz/blog/internal/services/author"
	"github.com/nnniyaz/blog/internal/services/bio"
	"github.com/nnniyaz/blog/internal/services/book"
	"github.com/nnniyaz/blog/internal/services/contact"
	"github.com/nnniyaz/blog/internal/services/project"
	"github.com/nnniyaz/blog/pkg/email"
)

type Service struct {
	Article articleService.ApplicationService
	Contact contactService.ContactService
	Author  authorService.AuthorService
	Bio     bioService.BioService
	Book    bookService.BookService
	Project projectService.ProjectService
}

func NewService(repos *repos.Repo, config *config.Config, emailService email.Email) *Service {
	return &Service{
		Article: articleService.NewArticleService(repos.RepoArticle),
		Contact: contactService.NewContactService(repos.RepoContact),
		Author:  authorService.NewAuthorService(repos.RepoAuthor),
		Bio:     bioService.NewBioService(repos.RepoBio),
		Book:    bookService.NewBookService(repos.RepoBook),
		Project: projectService.NewProjectService(repos.RepoProject),
	}
}