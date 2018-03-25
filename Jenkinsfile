stage('Pre Test'){
    echo 'Pulling Dependencies'
 
    sh 'go version'
    sh 'docker version'
    sh 'docker compose'
    sh 'go get -u github.com/golang/dep/dep'
    sh 'go get -u github.com/golang/lint/golint'
    sh 'go get github.com/tebeka/go2xunit'
    sh 'cd $GOPATH/src/project && dep ensure'
}
stage('Test'){
 
    //List all our project files with 'go list ./... | grep -v /vendor/ | grep -v github.com | grep -v golang.org'
    //Push our project files relative to ./src
    sh 'cd $GOPATH && go list ./... | grep -v /vendor/ | grep -v github.com | grep -v golang.org > projectPaths'
 
    //Print them with 'awk '$0="./src/"$0' projectPaths' in order to get full relative path to $GOPATH
    def paths = sh returnStdout: true, script: """awk '\$0="./src/"\$0' projectPaths"""
 
    echo 'Vetting'
 
    sh """cd $GOPATH && go tool vet ${paths}"""
 
    echo 'Linting'
    sh """cd $GOPATH && golint ${paths}"""
 
    echo 'Testing'
    sh """cd $GOPATH && go test -race -cover ${paths}"""
}
stage('Build'){
    echo 'Building Executable'
 
    //Produced binary is $GOPATH/src/project/project
    sh """cd $GOPATH/src/project/ && go build -ldflags '-s'"""
}
stage(' Publish'){
 
 
 
 
    //Find out commit hash
    sh 'git rev-parse HEAD > commit'
    def commit = readFile('commit').trim()
 
    //Find out current branch
    sh 'git name-rev --name-only HEAD > GIT_BRANCH'
    def branch = readFile('GIT_BRANCH').trim()
 
    //strip off repo-name/origin/ (optional)
    branch = branch.substring(branch.lastIndexOf('/') + 1)
 
    //def archive = "${GOPATH}/project-${branch}-${commit}.tar.gz"
 
    echo "Building Archive DockerImage"
   
    //sh """tar -cvzf ${archive} $GOPATH/src/project/project"""
 
    echo "Uploading ${archive} to  Downloads"
    sh """cd $GOPATH/src/project/ && docker-compose build"""
 
    // need to push the image to the docker hub
    //docker push registry-host:5000/predictive/project
    // Need to archive the image that has been created
 
    git add  ${archive}
    git commit -m "${branch}-${commit}"
    git push -u origin ${branch}
   
}
 
def notifyBuild(String buildStatus = 'STARTED') {
  // build status of null means successful
  buildStatus =  buildStatus ?: 'SUCCESSFUL'
 
  // Default values
  def colorName = 'RED'
  def colorCode = '#FF0000'
  def subject = "${buildStatus}: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'"
  def summary = "${subject} ${env.BUILD_URL}"
 
  // Override default values based on build status
  if (buildStatus == 'STARTED') {
    color = 'YELLOW'
    colorCode = '#FFFF00'
  } else if (buildStatus == 'SUCCESSFUL') {
    color = 'GREEN'
    colorCode = '#00FF00'
  } else {
    color = 'RED'
    colorCode = '#FF0000'
  }
 
  // Send notifications
  //slackSend (color: colorCode, message: summary)
}
