import requests

def fetch_github_user(username):
    url = f"https://api.github.com/users/{username}"
    response = requests.get(url)

    if response.status_code == 200:
        user_data = response.json()
        print(f"User: {user_data['login']}")
        print(f"Name: {user_data['name']}")
        print(f"Bio: {user_data['bio']}")
        print(f"Public Repos: {user_data['public_repos']}")
    else:
        print(f"Failed to fetch user: {response.status_code}")

if __name__ == "__main__":
    username = input("Enter a GitHub username: ")
    fetch_github_user(username)
