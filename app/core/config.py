from pydantic import BaseSettings


class Settings(BaseSettings):
    database_uri: str

    class Config:
        env_file = ".env"


settings = Settings()
