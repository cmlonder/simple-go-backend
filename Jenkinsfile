node {
    stage('Build') {
        echo "Workspace is : ${env.WORKSPACE}"
        sh '${env.WORKSPACE}/docker-compose up --build'
    }
}