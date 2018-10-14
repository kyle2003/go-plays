# pandora

Another self-educated and self-entertainmented project for learning go
Feel free to play 

# Design
This tool should provide the crawling in the background and a http service in the front for user to view pictures of 
website: xxgege.net

# Config
The config file will include 3 sections: [db], [download], [category], each sections will inlcude different attributes
please refer to:
- db_name : str database name is going to be created
- db_path : the path where to store the dbfile
- default_limit : Crawling limitation category specific
- image_path    : path for storing images

For the [category] section, you could specify the relative category you wanna to collect in key-value pairs

## Crawling process
Ideal processes
- Done => 1. Init Category
- Done => 2. Reap Subjects
- Done => 3. Reap Images
- Done => 4. Download and Update t_image (routine complete the job)
- 5. Provide web service to end User