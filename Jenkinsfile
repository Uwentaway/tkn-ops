pipeline {
  agent any
  stages {
    stage('Stage') {
      steps {
        checkout scm
        tektonCreateRaw input: 'deploy/test/pipeline.yaml', inputType: 'FILE', namespace: 'tekton-devops-pipeline'
      }
    }
  }
}