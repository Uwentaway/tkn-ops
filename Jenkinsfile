pipeline {
  agent any
  stages {
    stage('Stage') {
      steps {
        checkout scm
        tektonCreateRaw(inputType: 'FILE', input: 'deploy/test/pipeline.yaml')
      }
    }
  }
}