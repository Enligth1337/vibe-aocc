# Makefile для автоматизации Git workflow
.PHONY: ai-commit ai-branch ai-mr help

# Создание AI коммита с новой веткой
ai-commit:
	@echo "🤖 Создаем AI коммит..."
	@timestamp=$$(date +%Y%m%d_%H%M%S); \
	branch_name="feature/ai-update-$$timestamp"; \
	echo "📝 Создаем ветку: $$branch_name"; \
	git checkout -b "$$branch_name"; \
	git add .; \
	git commit -m "AI update $$timestamp"; \
	git push origin "$$branch_name"; \
	echo "✅ Ветка $$branch_name создана и запушена"; \
	echo "🔗 Создайте MR в dev через веб-интерфейс"

# Создание новой ветки для AI
ai-branch:
	@echo "🌿 Создаем новую AI ветку..."
	@timestamp=$$(date +%Y%m%d_%H%M%S); \
	branch_name="feature/ai-update-$$timestamp"; \
	echo "📝 Создаем ветку: $$branch_name"; \
	git checkout -b "$$branch_name"; \
	echo "✅ Ветка $$branch_name создана"

# Создание MR (только для GitLab)
ai-mr:
	@echo "🔀 Создаем Merge Request..."
	@current_branch=$$(git branch --show-current); \
	if [[ "$$current_branch" == feature/* ]]; then \
		echo "📤 Пушим ветку $$current_branch..."; \
		git push origin "$$current_branch"; \
		echo "✅ Ветка запушена, создайте MR в dev через веб-интерфейс"; \
	else \
		echo "❌ Текущая ветка не является feature веткой"; \
		echo "💡 Сначала выполните: make ai-branch"; \
	fi

# Помощь
help:
	@echo "🤖 AI Git Workflow Commands:"
	@echo ""
	@echo "  make ai-commit  - Создать коммит в новой ветке и запушнуть"
	@echo "  make ai-branch  - Создать новую ветку для AI изменений"
	@echo "  make ai-mr      - Запушнуть текущую ветку для создания MR"
	@echo "  make help       - Показать эту справку"
	@echo ""
	@echo "📋 Workflow:"
	@echo "  1. make ai-commit  (создает ветку + коммит + пуш)"
	@echo "  2. Создать MR в dev через веб-интерфейс"
	@echo "  3. После аппрува - мерж в dev"
