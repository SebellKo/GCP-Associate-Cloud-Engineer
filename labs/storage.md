# 실습: 비공개 Cloud Storage

[공통 변수](../00-getting-started/cost-safety.md)를 설정한다. bucket 이름은 전역 유일해야 한다. 아래 고정 suffix가 충돌하면 본인 전용 suffix를 바꾼다. 개인 데이터는 사용하지 않는다.

```bash
export BUCKET="${PROJECT_ID}-ace-lab-objects"
gcloud services enable storage.googleapis.com --project="$PROJECT_ID"
gcloud storage buckets create "gs://$BUCKET" --project="$PROJECT_ID"   --location="$REGION" --default-storage-class=STANDARD --uniform-bucket-level-access
gcloud storage buckets update "gs://$BUCKET" --public-access-prevention
printf 'ace storage lab\n' > /tmp/ace-object.txt
gcloud storage cp /tmp/ace-object.txt "gs://$BUCKET/hello.txt"
gcloud storage cp "gs://$BUCKET/hello.txt" /tmp/ace-object-downloaded.txt
cmp /tmp/ace-object.txt /tmp/ace-object-downloaded.txt
gcloud storage buckets get-iam-policy "gs://$BUCKET"
```

성공 기준: `cmp`가 종료 코드 0이고 IAM에 의도하지 않은 allUsers가 없는 것이다. 익명 공개를 하지 않고 인증된 작업을 사용한다.

## Lifecycle 설계

아래 예시는 30일 경과한 객체를 Nearline으로 바꾼다. 설명용 JSON이며 자동 적용하지 않는다. 최소 저장 기간과 조회 비용도 계산한다.

```json
{"rule":[{"action":{"type":"SetStorageClass","storageClass":"NEARLINE"},"condition":{"age":30,"matchesStorageClass":["STANDARD"]}}]}
```

## 정리

테스트 객체와 해당 버킷만 삭제한다.

```bash
gcloud storage rm "gs://$BUCKET/hello.txt"
gcloud storage buckets delete "gs://$BUCKET"
```

soft delete·versioning·retention 설정에 따라 보관·삭제 결과가 달라질 수 있다. 테스트 객체만 확인해 삭제하며 다른 버킷에는 실행하지 않는다.

[Cloud Storage docs](https://cloud.google.com/storage/docs/creating-buckets)
