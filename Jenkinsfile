pipeline {
  agent any
  stages {
    stage('get branch'){
      steps{
        script{
          print(branch)
        }
      }
    }
    stage('deploy to dev'){
        when{
          branch 'dev'
        }
        steps{
            tektonCreateRaw input: 'deploy/dev/pipeline.yaml', inputType: 'FILE', namespace: 'tekton-devops-pipeline'
        }
    }
    stage('deploy to test'){
        when{
          branch 'test'
        }
        steps{
            tektonCreateRaw input: 'deploy/test/pipeline.yaml', inputType: 'FILE', namespace: 'tekton-devops-pipeline'
        }
    }
    stage('deploy to uat'){
        when{
          branch 'uat'
        }
        steps{
            tektonCreateRaw input: 'deploy/uat/pipeline.yaml', inputType: 'FILE', namespace: 'tekton-devops-pipeline'
        }
    }
    stage('deploy to pre'){
        when{
          branch 'pre'
        }
        steps{
            tektonCreateRaw input: 'deploy/pre/pipeline.yaml', inputType: 'FILE', namespace: 'tekton-devops-pipeline'
        }
    }
    stage('deploy to prod'){
        when{
          branch 'prod'
        }
        steps{
            tektonCreateRaw input: 'deploy/prod/pipeline.yaml', inputType: 'FILE', namespace: 'tekton-devops-pipeline'
        }
    }
  }
}