node {
    stage('clean') {
        cleanWs()
    }

     stage('checkout') {
        checkout([
            $class: 'GitSCM',
            branches: scm.branches,
            extensions: scm.extensions +
                                            [[$class: 'CleanBeforeCheckout'],
                                            [$class: 'CheckoutOption', timeout: 60],
                                            [$class: 'CloneOption', depth: 0, noTags: false, reference: '', shallow: false, timeout: 60]],
            userRemoteConfigs: scm.userRemoteConfigs
         ])
     }

    stage('Start testing environment') {
        echo "Workspace is : ${env.WORKSPACE}"
        sh 'docker-compose up -d'
    }

    stage('Stop testing environment') {
        sh 'docker-compose down'
    }

    stage('get config file') {
            sh "wget https://raw.githubusercontent.com/Blazemeter/taurus/master/examples/jmeter/stepping.yml"
    }

    stage("run test") {
        bzt "stepping.yml"
    }
}