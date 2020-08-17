node {
    stage('clean') {
        cleanWs()
    }

     stage('checkout') {
        checkout scm
     }

    stage('Start testing environment') {
        echo "Workspace is : ${env.WORKSPACE}"
        sh 'docker-compose up -d'
    }

    stage('Stop testing environment') {
        sh 'docker-compose down'
    }

    stage('get config file') {
            sh 'wget "https://raw.githubusercontent.com/Blazemeter/taurus/master/examples/jmeter/stepping.yml"'
    }

    stage("run test") {
        bzt "stepping.yml"
    }
}