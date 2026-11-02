# Chapter 10. Cloud Storage·수명 주기·파일 스토리지

[← 이전 장](chapter-09.md) · [전체 목차](../README.md) · [다음 장 →](chapter-11.md)

공식 가이드 연결: **2.2 / 3.2** · 권장 2~3시간(통합 실습은 4~6시간)

## 1. 개념

### 객체·블록·파일을 구분한다
Cloud Storage는 bucket 안에 object를 저장한다. VM 블록 디스크나 공유 POSIX 파일 시스템과 다르다. 공유 NFS가 필요하면 Filestore, 엔터프라이즈 파일 요구에는 NetApp Volumes, HPC 병렬 파일 I/O에는 Managed Lustre 같은 선택지를 검토한다. 이름만 외우지 말고 접근 프로토콜과 작업 특성을 본다.

Standard는 자주 읽는 데이터, Nearline·Coldline·Archive는 드물게 읽는 데이터에 적합하다. 세 등급의 최소 저장 기간은 각각 30·90·365일이며 조회·조기 삭제 비용도 고려한다. Archive가 오프라인 테이프처럼 복구에 며칠 걸린다고 생각하지 않는다. 위치와 저장 등급은 서로 다른 설정이다.

Uniform bucket-level access는 객체 ACL 대신 IAM으로 권한을 일원화한다. Public access prevention은 공개 접근을 제한한다. Lifecycle은 조건에 맞춰 등급 변경·삭제를 수행하고 retention policy는 보존 기간 전에 삭제하지 못하도록 제한한다. Soft delete나 버전 보관은 비용·복구 요구와 함께 검토한다.

## 2. 개념 확인 문제

답을 먼저 적고 해설을 연다. 각 질문에 본인의 예시를 하나 붙인다.

1. 버킷 위치와 저장 등급은 같은가?

2. Coldline 최소 저장 기간은?

3. Lifecycle과 retention의 차이는?

## 3. 실습과 관찰

[버킷 실습](../labs/storage.md)을 수행한다. 비공개 Standard 버킷에 테스트 객체를 올리고 다운로드해 비교한다. 수명 주기 JSON을 읽어 30일 조건을 설명하되 삭제 규칙을 중요한 버킷에 적용하지 않는다. 객체·버킷·soft-deleted 데이터 보관 비용을 확인한다.

## 4. 실전 시험형 예제

학습용 자체 제작 문제다. 실제 출제 문항이나 시험 덤프가 아니다.

### 문제 1 — 단일 선택

매일 읽는 웹 이미지가 있다. 기본적으로 가장 적합한 등급은?

A. Nearline
B. Coldline
C. Archive
D. Standard

### 문제 2 — 단일 선택

여러 VM의 기존 앱이 공유 NFS 디렉터리를 요구한다.

A. Pub/Sub를 파일 시스템으로 쓴다
B. Cloud Storage를 블록 디스크로 간주한다
C. Cloud DNS에 파일을 저장한다
D. Filestore를 검토한다

### 문제 3 — 단일 선택

로그 객체를 일정 시간이 지나면 자동으로 저비용 등급으로 바꾸고 싶다.

A. 모든 객체를 공개한다
B. Object Lifecycle Management를 설정한다
C. 프로젝트 이름을 바꾼다
D. VM 자동 확장을 켠다

<details>
<summary>개념 확인 정답과 시험형 문제 해설</summary>

### 개념 확인

1. 아니다.

2. 90일이다.

3. 자동 관리 동작과 조기 삭제 제한이다.

### 시험형 예제

**1. D** — 빈번한 읽기에는 Standard가 적합하다. 나머지는 조회 비용·최소 기간을 고려하는 비빈번 접근 등급이다.

**2. D** — 프로토콜이 NFS이므로 관리형 파일 저장소가 맞다. 객체·메시지·DNS는 NFS 대체가 아니다.

**3. B** — 객체 조건별 등급 전환은 lifecycle 기능이다. 이름·공개화·VM 확장은 객체 정책이 아니다.

</details>

## 5. 완료 체크리스트

- [ ] 핵심 개념을 문서 없이 설명하고 내 서비스의 예를 들었다.
- [ ] 개념 확인 3문항과 시험형 3문항을 풀고 오답 이유를 기록했다.
- [ ] 실습 결과 또는 설계 산출물을 남겼다.
- [ ] 생성 자원의 삭제와 잔존 비용을 확인했다(조회·설계만 했다면 해당 없음).

## 공식 자료

- [관련 제품/학습 문서](https://cloud.google.com/storage/docs/storage-classes)
- [시험 범위와 최신성 관리](../00-getting-started/exam-overview.md)

[← 이전 장](chapter-09.md) · [다음 장 →](chapter-11.md)
