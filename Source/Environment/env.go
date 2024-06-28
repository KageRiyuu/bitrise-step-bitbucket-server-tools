package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	BITBUCKET_API_ACCESS_TOKEN string
	BITBUCKET_DOMAIN           string
	BITBUCKET_PROJECT_KEY      string
	BITBUCKET_REPO_SLUG        string

	BITRISE_API_ACCESS_TOKEN string
	BITRISE_APP_SLUG         string

	FUNCTION           string
	PR_ID              string
	BITRISE_BUILD_SLUG string
)

func LoadEnvironment() {
	fmt.Println("Loading env variables...")
	err := godotenv.Load()

	if err != nil {
		fmt.Printf("Env variable load error: %s", err)
		fmt.Println()
	}

	BITBUCKET_API_ACCESS_TOKEN = os.Getenv("bitbucket_api_access_token")
	BITBUCKET_DOMAIN = os.Getenv("bitbucket_domain")
	BITBUCKET_PROJECT_KEY = os.Getenv("bitbucket_project_key")
	BITBUCKET_REPO_SLUG = os.Getenv("bitbucket_repo_slug")

	BITRISE_API_ACCESS_TOKEN = os.Getenv("bitrise_api_access_token")
	BITRISE_APP_SLUG = os.Getenv("bitrise_app_slug")

	FUNCTION = os.Getenv("function")
	PR_ID = os.Getenv("pr_id")
	BITRISE_BUILD_SLUG = os.Getenv("bitrise_build_slug")

	DumpInputs()
	fmt.Println("Loading env finished !!")
}

func DumpInputs() {
	fmt.Println("-----------------------------------------")
	fmt.Println("Inputs: ")
	fmt.Println("Bitbucket Base Url: " + BITBUCKET_DOMAIN)
	fmt.Println("Bitbucket Project Key: " + BITBUCKET_PROJECT_KEY)
	fmt.Println("Bitbucket Repo Slug: " + BITBUCKET_REPO_SLUG)
	fmt.Printf("Bitbucket Token Provided: %t", len(BITBUCKET_API_ACCESS_TOKEN) != 0)
	fmt.Println()
	fmt.Println()

	fmt.Println("Bitrise App Slug: " + BITRISE_APP_SLUG)
	fmt.Printf("Bitrise Token Provided: %t", len(BITRISE_API_ACCESS_TOKEN) != 0)
	fmt.Println()
	fmt.Println()

	fmt.Println("Selected Function: " + FUNCTION)
	fmt.Println("Pull Request Id: " + PR_ID)
	fmt.Println("Bitrise Build Slug: " + BITRISE_BUILD_SLUG)
	fmt.Println("-----------------------------------------")
	fmt.Println("")
}

const (
	SKIPPED_VERIFICATION = "SKIPPED_VERIFICATION"
)

func DumpOutputs() {
	fmt.Println()
	fmt.Println("........................................")
	fmt.Println("Outputs: ")
	fmt.Println("Verification Skipped: " + os.Getenv("SKIPPED_VERIFICATION"))
}
