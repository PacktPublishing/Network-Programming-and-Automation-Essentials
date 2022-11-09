from loguru import logger


logger.add(
    "log-file-{time}.txt",
    rotation="1 MB",
    colorize=False,
    level="ERROR",
)

logger.debug("That's not going to show")
logger.warning("This will not show")
logger.error("Got an error")
