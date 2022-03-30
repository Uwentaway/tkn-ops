pipeline {
  agent any
  stages {
    stage('run tekton'){
        steps{
            tektonCreateRaw enableCatalog: true, input: 'deploy/test/pipeline.yaml', inputType: 'FILE', namespace: 'tekton-devops-pipeline'
        }
    }
  }
}