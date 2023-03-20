#!/bin/bash

folder_name="$1 at $(date)"
mkdir "$folder_name"

mkdir "$folder_name/about_me"
mkdir "$folder_name/about_me/personal"
echo "https://www.facebook.com/$2" > "$folder_name/about_me/personal/facebook.txt"
mkdir "$folder_name/about_me/professional"
echo "https://www.linkedin.com/in/$3" > "$folder_name/about_me/professional/linkedin.txt"

mkdir "$folder_name/my_friends"
curl -s "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt" > "$folder_name/my_friends/list_of_my_friends.txt"

mkdir "$folder_name/my_system_info"
echo "My username: $USER" > "$folder_name/my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "$folder_name/my_system_info/about_this_laptop.txt"

ping -c 3 google.com > "$folder_name/my_system_info/internet_connection.txt"

tree "$folder_name"