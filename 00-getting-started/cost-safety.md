# 실습 준비·비용·정리

## 준비

개인 학습 전용 프로젝트와 본인에게 사용 권한이 있는 결제 환경을 사용한다. 회사·운영 프로젝트에서는 이 교재의 변경·삭제 명령을 실행하지 않는다. Cloud Shell을 기본 터미널로 사용한다. 로컬 실습이라면 Google Cloud CLI, 필요한 Docker/Go/Terraform/kubectl을 각각 공식 절차로 설치한다.

```bash
export PROJECT_ID='본인의-학습-프로젝트-id'
export REGION='asia-northeast3'
export ZONE='asia-northeast3-a'
gcloud auth list
gcloud config set project "$PROJECT_ID"
gcloud projects describe "$PROJECT_ID"
```

샘플 위치에서 제품·머신·quota가 제공되는지 먼저 확인한다. 실습마다 새 터미널을 쓰면 변수를 다시 설정한다. `PROJECT_ID`는 교재 문구 그대로 쓰지 않는다. 생성 권한·조직 정책에 의해 일부 단계가 제한될 수 있다.

## 비용 원칙

alerts-only 예산은 비용 자동 차단이 아니다. 별도 spend cap 기능은 지원 서비스와 제한을 확인해야 하며 전 서비스 하드 캡으로 가정하지 않는다. 무료 체험·무료 한도는 자격·위치·상품·기간에 따라 다르므로 [현재 무료 사용 정책](https://cloud.google.com/free/docs/free-cloud-features)과 [가격 계산기](https://cloud.google.com/products/calculator)를 확인한다. 무료라고 단정한 고정 금액은 제시하지 않는다.

| 실습 | 주요 비용 요소 | 기본 진행 방식 |
|---|---|---|
| VM | 실행 시간·디스크·IP·snapshot | 짧게 실행 후 삭제 |
| Storage | 용량·요청·전송·soft delete | 작은 Standard 테스트 객체 |
| GKE | cluster·Pod/노드·디스크·네트워크 | 선택 유료 실습 |
| DB | 인스턴스·스토리지·백업·복제본 | 설계 먼저, 선택 유료 실습 |
| LB/NAT/VPN | 시간·처리·전송·IP | 설계 먼저 |
| Cloud Run 빌드 | 실행·build·registry·로그 | 짧은 통합 실습 |
| GPU/TPU/AI runtime | 가속기·세션·스토리지 | 생성 없이 개념 확인 |

## 삭제 점검

애플리케이션 삭제 후 연결 자원을 확인한다. VM→디스크·snapshot·예약 IP, GKE→LB·디스크, Run→registry 이미지·build 저장소, DB→복원본·backup 등이다. 삭제 전 `describe/list`로 이름과 프로젝트를 재확인한다. 저장소에는 운영 프로젝트 전체 삭제나 전 자원 일괄 파괴 스크립트를 제공하지 않는다.

실습별 명시 자원 삭제 후 Billing Reports에서 잔존 비용을 확인한다. 보고에는 지연이 있어 즉시 0 표시만으로 완료를 판단하지 않는다. 학습 자료에는 토큰·키·결제 정보·개인 데이터 대신 자원 이름과 관찰 결과만 기록한다.
