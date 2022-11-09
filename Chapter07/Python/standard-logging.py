import logging

logging.basicConfig(
    filename="log-file.txt",
    level=logging.ERROR,
    format="%(asctime)s.%(msecs)03d %(levelname)s: %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S",
)
logging.debug("This won't show, level is set to info")
logging.info("Info is not that important as well")
logging.warning("Warning will not show as well")
logging.error("This is an error")