### English

- Use Qoder's real per-model context windows instead of a hardcoded 180000. Qoder reports a `default_context_window` and a set of `available_context_windows` per model, but only `max_input_tokens` was kept, so every Qoder model showed the same 180000 window — under-sizing glm-5.3-flash (1M) and over-sizing deepseek-v4-pro (96K). The worker now forwards the window metadata and the console defaults each model to the window Qoder actually reports, with the largest selectable window exposed as the model's max tier. Falls back to the previous default only when Qoder reports no window.

### 中文

- Qoder 模型改用上游真实上下文窗口，不再一律硬编码 180000。Qoder 每个模型都会上报 `default_context_window` 与一组 `available_context_windows`，但此前只保留了 `max_input_tokens`，导致所有 Qoder 模型都显示同一个 180000 窗口——把 glm-5.3-flash（1M）压小、把 deepseek-v4-pro（96K）放大。现在 worker 透传窗口元数据，控制台按 Qoder 实际上报的窗口作为每个模型的默认值，并把可选窗口里的最大值作为该模型的 Max 档。仅当上游未上报窗口时才回退到旧的默认值。
