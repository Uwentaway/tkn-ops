pipeline {
  agent any
  stages {
    stage('checkout') {
      steps {
        checkout scm
      }
    }
    stage('run tekton'){
        steps{
            tektonCreateRaw enableCatalog: true, input: 'deploy/test/pipeline.yaml', inputType: 'FILE', namespace: 'tekton-devops-pipeline'
        }
    }
  }
}