# Chapter 04. Cloud Shell·gcloud·인증과 구성

[← 이전 장](chapter-03.md) · [전체 목차](../README.md) · [다음 장 →](chapter-05.md)

공식 가이드 연결: **1.1 / 4.2** · 권장 2~3시간(통합 실습은 4~6시간)

## 1. 개념

### 명령이 실행되는 대상을 고정한다
Cloud Shell은 브라우저 기반 터미널이며 사용자가 배포하는 VM과 다르다. `gcloud`는 Google Cloud API를 호출하는 CLI다. 로그인 계정, 프로젝트, 기본 리전·존을 확인한 뒤 명령을 실행한다. 여러 환경은 named configuration으로 나누고 위험한 변경에는 `--project`를 명시한다.

`gcloud auth login`은 CLI 사용자 인증을 설정한다. Application Default Credentials(ADC)는 앱과 Terraform 같은 라이브러리가 기본 자격 증명을 찾는 방식이다. 로컬 ADC 설정은 `gcloud auth application-default login`으로 할 수 있다. CLI 로그인이 곧 모든 SDK의 ADC 설정을 뜻하지는 않는다. Google Cloud 안의 앱은 연결된 서비스 계정과 메타데이터 기반 인증을 우선 사용한다.

리소스 이름과 위치가 맞지 않으면 존재하는 VM도 찾지 못한다. 변경 전에 list/describe로 대상을 확인하고 JSON 출력이나 `--format`으로 필요한 필드만 본다.

## 2. 개념 확인 문제

답을 먼저 적고 해설을 연다. 각 질문에 본인의 예시를 하나 붙인다.

1. gcloud 로그인과 ADC는 같은 설정인가?

2. --project를 명시하는 이유는?

3. 존 오류와 IAM 오류는 어떻게 구분하는가?

## 3. 실습과 관찰

```bash
gcloud auth list
gcloud config list
gcloud config configurations list
gcloud projects describe "$PROJECT_ID" --format='value(projectNumber)'
```
원하는 계정과 프로젝트인지 확인한다. 이 명령들은 조회용이며 생성 리소스는 없다. 로컬 ADC 파일이나 액세스 토큰을 저장소에 넣지 않는다.

## 4. 실전 시험형 예제

학습용 자체 제작 문제다. 실제 출제 문항이나 시험 덤프가 아니다.

### 문제 1 — 단일 선택

개발·운영 프로젝트를 번갈아 관리한다. 실수를 줄이는 방법은?

A. 모든 계정을 Owner로 바꾼다
B. 환경별 named configuration과 명시적 --project를 사용한다
C. 프로젝트 이름만 기억한다
D. 명령 기록을 지운다

### 문제 2 — 단일 선택

로컬 gcloud 조회는 되지만 ADC를 사용하는 앱에서 기본 자격 증명을 못 찾는다.

A. VM을 Spot으로 바꾼다
B. 앱의 ADC 탐색 경로와 로컬 ADC 설정을 확인한다
C. 버킷을 공개한다
D. DNS를 삭제한다

### 문제 3 — 단일 선택

VM 조회가 다른 존에서 실패한다. 가장 적절한 조치는?

A. 조직 정책을 전부 해제한다
B. 새 결제 계정을 만든다
C. 메모리를 두 배로 늘린다
D. VM 목록으로 실제 존을 확인하고 해당 존을 지정한다

<details>
<summary>개념 확인 정답과 시험형 문제 해설</summary>

### 개념 확인

1. 용도와 저장된 자격 증명이 구분된다.

2. 잘못된 기본 프로젝트에 변경하는 일을 줄인다.

3. 위치·존재 여부와 실제 오류 코드·권한을 각각 확인한다.

### 시험형 예제

**1. B** — 실행 컨텍스트를 분리하고 대상을 명시해야 한다. 기록 삭제나 권한 확대, 기억 의존은 대상 혼동을 줄이지 못한다.

**2. B** — CLI와 앱 자격 증명 경로가 다를 수 있다. 공개화·VM 정책·DNS는 ADC 탐색을 해결하지 않는다.

**3. D** — 자원의 위치를 먼저 맞춘다. 정책 해제·결제·사양 증설은 위치 불일치 해결책이 아니다.

</details>

## 5. 완료 체크리스트

- [ ] 핵심 개념을 문서 없이 설명하고 내 서비스의 예를 들었다.
- [ ] 개념 확인 3문항과 시험형 3문항을 풀고 오답 이유를 기록했다.
- [ ] 실습 결과 또는 설계 산출물을 남겼다.
- [ ] 생성 자원의 삭제와 잔존 비용을 확인했다(조회·설계만 했다면 해당 없음).

## 공식 자료

- [관련 제품/학습 문서](https://cloud.google.com/sdk/docs/configurations)
- [시험 범위와 최신성 관리](../00-getting-started/exam-overview.md)

[← 이전 장](chapter-03.md) · [다음 장 →](chapter-05.md)
