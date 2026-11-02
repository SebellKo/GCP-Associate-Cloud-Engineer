# 실습: Terraform으로 비공개 bucket 관리

Terraform CLI와 ADC가 필요하다. Cloud Shell의 환경 인증을 사용하거나 로컬에서 `gcloud auth application-default login`을 설정한다. 프로젝트 API·권한·결제 상태를 먼저 확인한다. 아래 bucket은 작은 테스트 자원이며 저장·요청·전송·보관 비용이 있을 수 있다.

```bash
cd labs/terraform
terraform init
terraform fmt -check
terraform validate
terraform plan -var="project_id=$PROJECT_ID" -var="region=$REGION"   -var="bucket_name=${PROJECT_ID}-ace-tf-lab" -out=ace.tfplan
terraform show ace.tfplan
```

계획이 테스트 bucket 하나만 만드는지 확인하고 허용할 때 `terraform apply ace.tfplan`을 실행한다. `.terraform.lock.hcl`은 provider 선택 재현을 위해 검토 후 관리한다. state/plan은 Git에 넣지 않는다. 이 실습은 개인 로컬 state 예제이며 팀 작업은 접근 제어·잠금이 있는 원격 backend로 확장한다.

## 정리

```bash
terraform plan -destroy -var="project_id=$PROJECT_ID" -var="region=$REGION"   -var="bucket_name=${PROJECT_ID}-ace-tf-lab" -out=destroy.tfplan
terraform show destroy.tfplan
terraform apply destroy.tfplan
```

bucket에 임의 객체를 넣었다면 먼저 해당 테스트 객체만 삭제한다. force_destroy는 기본으로 켜지 않았다. state를 삭제해도 Cloud 자원이 삭제되는 것은 아니다. 종료 후 Cloud Storage 목록에서 확인한다.

[공식 Terraform 기본 명령](https://cloud.google.com/docs/terraform/basic-commands)
