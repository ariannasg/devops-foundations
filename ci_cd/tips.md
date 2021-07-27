* Add a prefix to the version of the artifact so we can increase that number in case we need to change
artifact repository platform. And use the version of the build in the version of the artifact, e.g.:
  `version = 1.$BUILD_NUMBER` ($BUILD_NUMBER is a Jenkins var in this case)
  
* Add tags to the code with the build number (add extra step on pipeline to do this)
