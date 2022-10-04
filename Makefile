COLOR_YELLOW="\033[33m"
COLOR_GREEN="\033[32m"
COLOR_BLUE="\033[36m"
COLOR_WHITE="\033[37m"

run:
	@echo ${COLOR_YELLOW}Running Application...${COLOR_WHITE}
	@go run main.go

install:
	@echo ${COLOR_YELLOW}Running Tidy...${COLOR_WHITE}
	@go mod tidy
	@echo "--------------------------------------------"
	@echo ${COLOR_YELLOW}Running Vendor...${COLOR_WHITE}
	@go mod vendor
	@echo "--------------------------------------------"
	@echo ${COLOR_GREEN} DONE

container:
	@echo ${COLOR_YELLOW}Starting container...${COLOR_WHITE}
	@docker-compose up -d

kill-container:
	@echo ${COLOR_YELLOW}Killing container...${COLOR_WHITE}
	@docker-compose down
	@echo "--------------------------------------------"
	@echo ${COLOR_YELLOW}Prunning containers...${COLOR_WHITE}
	@docker container prune -f
	@echo "--------------------------------------------"
	@echo ${COLOR_YELLOW}Prunning volumes...${COLOR_WHITE}
	@docker volume prune -f
	@echo "--------------------------------------------"
	@echo ${COLOR_GREEN} DONE