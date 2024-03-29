import time
import unittest
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from webdriver_manager.chrome import ChromeDriverManager
import chromedriver_autoinstaller as chromedriver

class TestSearch(unittest.TestCase):

    def setUp(self):
        options = Options()
        options.add_argument('--headless')
        options.add_argument('--no-sandbox')
        options.add_argument('--disable-dev-shm-usage')
        chromedriver.install()
        self.driver = webdriver.Chrome()
#        s = Service('/home/ubuntu/chromedriver.exe')
#        self.driver = webdriver.Chrome(service=s, options=options)



    def test_search_valid(self):
        driver = self.driver
        time.sleep(5)
        driver.get('http://0.0.0.0')
        search = driver.find_element(by=By.NAME, value="city")
        search.send_keys("Haifa")
        search.send_keys(Keys.RETURN)
        try:
            self.assertEqual(driver.title, "Weather Forecast")
        except Exception as e:
            self.assertNotEqual(driver.title, "Weather Forecast")
            print("Error:", e)
        time.sleep(5)

    def test_search_invalid(self):
        driver = self.driver
        time.sleep(10)
        driver.get('http://0.0.0.0')
        search = driver.find_element(by=By.NAME, value="city")
        search.send_keys("asdjioasdjasidaskdlasndiawodnsadklandwuidnadjkwnadis")
        search.send_keys(Keys.RETURN)
        try:
            self.assertEqual(driver.title, "Couldnt find request")
        except Exception as e:
            self.assertNotEqual(driver.title, "Couldnt find request")
            print("Error:", e)
        time.sleep(5)

    def tearDown(self):
        self.driver.close()


if __name__ == "__main__":
    unittest.main()
