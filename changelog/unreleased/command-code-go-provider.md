### English

- Add an experimental **Command Code** (`commandcode.ai`) provider for the **$1 Go plan**. Go accounts have no `/provider/v1/*` API access (those paths return `403 upgrade_required`), so the adapter speaks `/alpha/generate` — the undocumented envelope the `cmd` CLI itself uses on every turn, which is not plan-gated and serves the whole model catalog. Auth is a single `user_…` key, pasted through the existing PAT tab; the model list is read from the anonymous `/provider/v1/models`, and account credits are shown from `/alpha/billing/credits`. Streaming is rewritten from the upstream newline-delimited JSON into OpenAI SSE. Experimental: `/alpha/generate` is unpublished and version-coupled, so the pinned CLI version fails loudly rather than degrading silently, and the provider is not production-ready.

### 中文

- 新增实验性 **Command Code**（`commandcode.ai`）渠道，面向 **$1 Go 套餐**。Go 档没有 `/provider/v1/*` API 访问权限（该路径返回 `403 upgrade_required`），因此适配层走 `/alpha/generate` —— 这是 `cmd` CLI 每次调用都用的未公开 envelope，不受档位限制，且覆盖整个模型目录。认证只需一把 `user_…` 密钥，经现有 PAT 页签粘贴；模型列表读取匿名的 `/provider/v1/models`；额度从 `/alpha/billing/credits` 展示。流式响应由上游的换行分隔 JSON 改写为 OpenAI SSE。实验性：`/alpha/generate` 未公开且与 CLI 版本耦合，故钉版不匹配时明确报错而非静默降级，不承诺生产可用。
