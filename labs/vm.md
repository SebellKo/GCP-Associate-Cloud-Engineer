# 실습: 외부 IP 없는 VM과 IAP

전제: Compute API 활성화 권한, VM/네트워크 생성 권한, IAP tunnel 접근 권한, OS Login 권한(관리 작업에는 Admin Login), 필요한 경우 VM 서비스 계정 actAs 권한. 개인 학습 계정의 권한도 정책에 따라 다르다. [준비 변수](../00-getting-started/cost-safety.md)를 먼저 설정한다.

## 생성

```bash
gcloud services enable compute.googleapis.com iap.googleapis.com --project="$PROJECT_ID"
gcloud compute networks create ace-lab-vpc --subnet-mode=custom --project="$PROJECT_ID"
gcloud compute networks subnets create ace-lab-subnet   --network=ace-lab-vpc --region="$REGION" --range=10.42.0.0/24 --project="$PROJECT_ID"
gcloud compute firewall-rules create ace-lab-iap-ssh   --network=ace-lab-vpc --direction=INGRESS --action=ALLOW   --rules=tcp:22 --source-ranges=35.235.240.0/20 --target-tags=ace-lab-ssh --project="$PROJECT_ID"
gcloud compute instances create ace-lab-vm   --project="$PROJECT_ID" --zone="$ZONE" --machine-type=e2-micro   --subnet=ace-lab-subnet --no-address --tags=ace-lab-ssh   --image-family=debian-12 --image-project=debian-cloud   --metadata=enable-oslogin=TRUE --no-service-account --no-scopes
gcloud compute ssh ace-lab-vm --tunnel-through-iap --zone="$ZONE" --project="$PROJECT_ID"
```

VM 안에서 `hostname`, `uname -a`, `df -h`를 실행하고 `exit`한다. 외부 IP가 없음을 describe로 확인한다. NAT가 없으므로 인터넷 패키지 다운로드가 안 될 수 있으며 SSH 성공과 모순되지 않는다.

실패 시 IAP IAM→OS Login→방화벽 target/source→리전/존→조직 정책 순으로 확인한다. 모든 ingress를 열지 않는다. VM Manager를 확장 실습하려면 OS Config API·agent·실행 계정 등 별도 공식 전제를 확인한다.

## 정리

```bash
gcloud compute instances delete ace-lab-vm --zone="$ZONE" --project="$PROJECT_ID"
gcloud compute firewall-rules delete ace-lab-iap-ssh --project="$PROJECT_ID"
gcloud compute networks subnets delete ace-lab-subnet --region="$REGION" --project="$PROJECT_ID"
gcloud compute networks delete ace-lab-vpc --project="$PROJECT_ID"
gcloud compute disks list --project="$PROJECT_ID"
gcloud compute snapshots list --project="$PROJECT_ID"
```

snapshot·복원 disk를 추가로 만들었다면 이름을 확인해 별도 삭제한다. 명령은 확인 프롬프트를 유지한다.

[IAP 문서](https://cloud.google.com/iap/docs/using-tcp-forwarding) · [OS Login](https://cloud.google.com/compute/docs/oslogin)
