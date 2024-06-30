# Bitbucket And Bitrise Tools

This step allows you to perform different functions on Services like Bitbucket Server, Bitrise, etc..
With time more function will be added to same repo.

## Common Inputs

Bitbucket Inputs:
- bitbucket_api_access_token
- bitbucket_domain
- bitbucket_project_key
- bitbucket_repo_slug

Bitrise Inputs:
- bitrise_api_access_token
- bitrise_app_slug

Function Inputs:
- function
- All params required by function. Check functions details below.

## Available Functions

1. Skip Verification
2. Check Pull Requests Deadline

# Skip Verification
This function will check if the PR Title contains "[SV]" tag, it will stop the build on bit rise and add the commit status "FAILED" on the head commit of the PR.

## Required Inputs
- pr_id:                    Bitbucket PR id, e.g: $BITRISE_PULL_REQUEST.
- bitrise_build_slug:       Build id on bitrise, e.g: $BITRISE_BUILD_SLUG

## Outputs
- SKIPPED_VERIFICATION:        True if found [SV] tag and aborted build, otherwise False.

# Check Pull Requests Deadline
This function will check all unapporved open Pull Requests and prepare a Json data `PULL_REQUESTS_DEADLINE`. It also checks tag [Deadline dd/mm] in Pull Request title and create a boolean flag `PULL_REQUESTS_DEADLINE_NEAR`. 

Pull Request title example:
[Deadline: 1/7][PROJ-12345][BUG][Deep Linking] - Web view 02 page can not be loaded completely and once close.
You can use this Json Data and Flag to perform operation e.g Notify Members to prioritise tasks accordingly.

## Required Inputs
- emails:                       List of pull request Authors email you what to check
- required_approval_count:      Minimum number of approvals required for a Pull Request
- title:                        Regular expression to filter Pull Request with Title. Check `How To Apply Regex` below.
- pr_id:                        Regular expression to filter Pull Request with Id. Check `How To Apply Regex` below.

## Outputs
- PULL_REQUESTS_DEADLINE_NEAR:      TRUE/FALSE if [Deadline dd/mm] is near 
- PULL_REQUESTS_DEADLINE:           Json String check `Sample JSON PULL_REQUESTS_DEADLINE`

### Sample JSON `PULL_REQUESTS_DEADLINE`
```
{
    "areDeadlinesNear": true,
    "pullRequests": [
        {
            "title": "[Deadline: 1/7][PRJ-12345] - Web view 02 page can not be loaded completely and once close it by tapping back, will open again automatically when tap RUNA banner under MEN gender",
            "url": "https://xxxxxxxm/projects/xxxxxx/repos/xxxxxxxxx/pull-requests/1403",
            "isDeadlineNear": true
        },
        {
            "title": "[Deadline: 15/7][PRJ-2569] - Will transfer to native search result by tapping RUNA when RUNA link format",
            "url": "https://xxxxxxxm/projects/xxxxxx/repos/xxxxxxxxx/pull-requests/1403",
            "isDeadlineNear": true
        }
    ]
}
```

# How To Apply Regex

Yoy can check and validate regular expression that can be used in Golang language.
https://regex101.com

To negate the provided reqular expression add "NRR " in begining of regular expression.

Example:
String Start with Ryu: "^Ryu"
String `DOES NOT` Start with Ryu: "NRR ^Ryu"

# How to use this Step

Add this in your bitrise.yml file and replace proper variables:

```
- git::https://github.com/KageRiyuu/bitrise-step-bitbucket-server-tools.git@main:
    title: Bitbucket Server Tools
    inputs:
    - pr_id: $BITRISE_PULL_REQUEST
```
Check above for all needed inputs