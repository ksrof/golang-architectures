// layered-gin/entity/validations.go

package entity

import (
	"html"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Prepare() cleans struct fields
func Prepare(character Character) {
	character.Name = html.EscapeString(character.Name)
	character.Title = html.EscapeString(character.Title)
	character.Gender = html.EscapeString(strings.TrimSpace(character.Gender))
	character.Race = html.EscapeString(character.Race)
	character.Class = html.EscapeString(character.Class)
	character.Affiliation = html.EscapeString(character.Affiliation)
	character.Status = html.EscapeString(strings.TrimSpace(character.Status))
}

// Validate() validates struct fields
func Validate(action string, character Character) error {
	validate := validator.New()

	switch strings.ToLower(action) {
	case "create":
		errs := validate.Var(character.Name, "required,min=3,max=255")
		if errs != nil {
			log.Printf("unable to validate name")
			return errs
		}
		
		errs = validate.Var(character.Title, "required,min=3,max=255")
		if errs != nil {
			log.Printf("unable to validate title")
			return errs
		}

		errs = validate.Var(character.Gender, "required,min=3,max=100")
		if errs != nil {
			log.Printf("unable to validate gender")
			return errs
		}


		errs = validate.Var(character.Race, "required,min=3,max=100")
		if errs != nil {
			log.Printf("unable to validate race")
			return errs
		}

		errs = validate.Var(character.Class, "required,min=3,max=100")
		if errs != nil {
			log.Printf("unable to validate class")
			return errs
		}

		errs = validate.Var(character.Affiliation, "required,min=3,max=100")
		if errs != nil {
			log.Printf("unable to validate affiliation")
			return errs
		}

		errs = validate.Var(character.Status, "required,min=3,max=50")
		if errs != nil {
			log.Printf("unable to validate status")
			return errs
		}

		return nil
	default:
		errs := validate.Var(character.Name, "required,min=3,max=255")
		if errs != nil {
			log.Printf("unable to validate name")
			return errs
		}
		
		errs = validate.Var(character.Title, "required,min=3,max=255")
		if errs != nil {
			log.Printf("unable to validate title")
			return errs
		}

		errs = validate.Var(character.Gender, "required,min=3,max=100")
		if errs != nil {
			log.Printf("unable to validate gender")
			return errs
		}


		errs = validate.Var(character.Race, "required,min=3,max=100")
		if errs != nil {
			log.Printf("unable to validate race")
			return errs
		}

		errs = validate.Var(character.Class, "required,min=3,max=100")
		if errs != nil {
			log.Printf("unable to validate class")
			return errs
		}

		errs = validate.Var(character.Affiliation, "required,min=3,max=100")
		if errs != nil {
			log.Printf("unable to validate affiliation")
			return errs
		}

		errs = validate.Var(character.Status, "required,min=3,max=50")
		if errs != nil {
			log.Printf("unable to validate status")
			return errs
		}

		return nil
	}
}