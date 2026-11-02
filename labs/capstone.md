# 통합 실습: 작은 Go API 배포·운영

[비용·공통 변수](../00-getting-started/cost-safety.md)를 확인한다. 저장소 루트에서 실행한다. Go를 처음부터 공부할 필요 없이 HTTP 서버의 PORT 처리와 배포 흐름에 집중한다.

## 1. 준비와 권한

배포 사용자에는 API 활성화, Artifact Registry 관리·업로드, Cloud Build 실행, Cloud Run 관리, runtime service account actAs와 필요한 IAM 변경 권한이 필요하다. 조직 환경에서는 관리자에게 학습 범위의 역할만 요청한다. 프로젝트 Owner를 기본 실습 역할로 권장하지 않는다.

Cloud Build는 프로젝트 설정에 따라 기본 실행 identity가 다를 수 있다. 빌드 시작 전에 Cloud Build 설정에서 실제 계정을 확인하고 Artifact Registry Writer와 필요한 로그 기록 권한을 부여한다. identity를 추측하지 않는다. API 활성화 직후 service agent 준비에 시간이 걸릴 수 있다.

```bash
gcloud services enable run.googleapis.com cloudbuild.googleapis.com artifactregistry.googleapis.com   iam.googleapis.com storage.googleapis.com --project="$PROJECT_ID"
export REPO='ace-lab-images'
export SERVICE='ace-lab-api'
export IMAGE="$REGION-docker.pkg.dev/$PROJECT_ID/$REPO/api:v1"
export RUNTIME_SA="ace-lab-runtime@$PROJECT_ID.iam.gserviceaccount.com"
gcloud artifacts repositories create "$REPO" --repository-format=docker   --location="$REGION" --project="$PROJECT_ID"
gcloud iam service-accounts create ace-lab-runtime --project="$PROJECT_ID"
```

## 2. 빌드와 비공개 배포

```bash
gcloud builds submit labs/app --tag="$IMAGE" --project="$PROJECT_ID"
gcloud run deploy "$SERVICE" --image="$IMAGE" --region="$REGION"   --project="$PROJECT_ID" --service-account="$RUNTIME_SA"   --no-allow-unauthenticated --min-instances=0 --max-instances=2
```

배포 사용자가 호출 권한까지 갖는다고 가정하지 않는다. 본인의 활성 사용자 계정을 확인하고 해당 서비스 범위에 `roles/run.invoker`를 부여한다. 아래는 사용자 계정 예시이며 federated principal이면 해당 principal 형식을 사용한다.

```bash
gcloud auth list
export CALLER_EMAIL='본인의-Google-계정-이메일'
gcloud run services add-iam-policy-binding "$SERVICE" --region="$REGION"   --project="$PROJECT_ID" --member="user:$CALLER_EMAIL" --role=roles/run.invoker
export SERVICE_URL="$(gcloud run services describe "$SERVICE" --region="$REGION"   --project="$PROJECT_ID" --format='value(status.url)')"
curl -fsS -H "Authorization: Bearer $(gcloud auth print-identity-token)" "$SERVICE_URL"
```

토큰을 문서·스크린샷에 저장하지 않는다. 성공 응답은 JSON의 `status: ok`다. 403이면 caller/Invoker/ID token을 보고 컨테이너 시작 오류면 PORT·로그·이미지를 본다.

## 3. 최소 권한과 관측

[Storage 실습](storage.md)에서 만든 bucket이 있다면 다음처럼 runtime 계정에 그 bucket만 읽게 설정한다.

```bash
gcloud storage buckets add-iam-policy-binding "gs://$BUCKET"   --member="serviceAccount:$RUNTIME_SA" --role=roles/storage.objectViewer
gcloud logging read 'resource.type="cloud_run_revision"'   --project="$PROJECT_ID" --limit=20
gcloud run revisions list --service="$SERVICE" --region="$REGION" --project="$PROJECT_ID"
```

샘플 앱은 HTTP 상태만 반환한다. bucket 읽기 권한 설정과 실제 SDK 호출은 구분한다. 확장 과제는 Storage SDK로 `hello.txt`를 읽고, 대상 bucket만 허용했는지 성공/실패를 각각 확인하는 것이다.

새 버전은 이미지 태그·digest를 기록하고 `--no-traffic` 배포 후 검증한다. `gcloud run services update-traffic`의 `--to-revisions`에 실제 revision 이름과 비율을 지정한다. 이전 revision 100% 복원 전 DB 호환성도 확인한다.

## 4. 정리

```bash
gcloud run services delete "$SERVICE" --region="$REGION" --project="$PROJECT_ID"
gcloud artifacts repositories delete "$REPO" --location="$REGION" --project="$PROJECT_ID"
```

Storage 실습을 수행했다면 해당 bucket binding을 제거하거나 테스트 bucket을 정리한 뒤 runtime 계정을 삭제한다.

```bash
gcloud iam service-accounts delete "$RUNTIME_SA" --project="$PROJECT_ID"
gcloud storage buckets list --project="$PROJECT_ID"
gcloud artifacts repositories list --location="$REGION" --project="$PROJECT_ID"
```

Cloud Build의 소스 업로드 bucket·logs bucket이 생성됐는지 콘솔에서 확인한다. 다른 작업과 공유하는 bucket은 삭제하지 말고 이번 테스트 객체만 확인해 정리한다. 계정 삭제만으로 다른 resource의 IAM binding이 모두 정리된다고 가정하지 않는다.
