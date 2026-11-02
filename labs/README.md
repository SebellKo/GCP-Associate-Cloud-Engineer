# 실습 목록

먼저 [비용·준비](../00-getting-started/cost-safety.md)를 읽고 PROJECT_ID/REGION/ZONE을 설정한다. 실습은 자동 실행되지 않는다. 같은 이름의 기존 자원이 있다면 다른 학습 전용 이름으로 바꿔 진행하고 정리 명령도 맞춘다.

| 실습 | 관련 장 | 성공 증거 |
|---|---|---|
| [VM/IAP](vm.md) | 07~09, 13 | private VM 내부 hostname |
| [비공개 bucket](storage.md) | 10, 24 | 업로드/다운로드 비교·IAM 범위 |
| [GKE](gke.md) | 19 | rollout·Pod·logs·port-forward |
| [Terraform](terraform.md) | 20 | 검토한 plan과 destroy 확인 |
| [통합 Cloud Run](capstone.md) | 16~18, 21~27 | 인증 호출·로그·revision·정리 |

조직 기능·네트워크 연결·고비용 DB·GPU/TPU·AI runtime은 해당 챕터의 설계 실습을 기본으로 한다. 여기에 실제 실행 기록이 없다면 GCP 동작 검증 완료로 간주하지 않는다.
