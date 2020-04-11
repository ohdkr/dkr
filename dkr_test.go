package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"os"
	"os/exec"
)

var _ = Describe("Dkr", func() {
	var session *gexec.Session
	var mainPath string

	var callDkrC = func(action string) *gexec.Session {
		pwd, _ := os.Getwd()

		composeFilenameArg := pwd + "/docker-compose-test-dummy.yml"
		args := []string{"c", "-f", composeFilenameArg, action}

		if action == "up" {
			args = append(args, "-d")
		}
		return runMain(mainPath, args)
	}

	BeforeEach(func() {
		mainPath = buildMain()
	})

	AfterEach(func() {
		gexec.CleanupBuildArtifacts()
	})

	Context("Helpers", func() {
		It("dkr exits with status code 0", func() {
			session = runMain(mainPath, []string{})
			session.Wait(10000)
			Eventually(session).Should(gexec.Exit(0))
		})

		It("calls dkr with no command", func() {
			session = runMain(mainPath, []string{})
			session.Wait(10000)
			Eventually(session).Should(gbytes.Say("Calling docker with \\[\\]"))
			Eventually(session).Should(gexec.Exit(0))
		})

		It("calls dkr --version", func() {
			session = runMain(mainPath, []string{"--version"})
			session.Wait(10000)
			Eventually(session).Should(gbytes.Say("Dkr version"))
			Eventually(session).Should(gbytes.Say("Docker version"))
			Eventually(session).Should(gbytes.Say("docker-compose version"))
			Eventually(session).Should(gexec.Exit(0))
		})

		It("calls dkr --help", func() {
			session = runMain(mainPath, []string{"--help"})
			session.Wait()
			Eventually(session).Should(gbytes.Say("aliases"))
			Eventually(session).Should(gbytes.Say("sh"))
			Eventually(session).Should(gbytes.Say("bash"))
			Eventually(session).Should(gbytes.Say("killall"))
			Eventually(session).Should(gbytes.Say("cleanup"))
			Eventually(session).Should(gbytes.Say("nuke"))
		})
	})

	Context("Docker/docker-compose proxy commands", func() {
		It("calls dkr c up", func() {
			session = callDkrC("up")
			session.Wait(60000)
			Eventually(session).Should(gexec.Exit(0))
		})

		It("checks if dkr c up worked", func() {
			session = runMain(mainPath, []string{"ps"})
			session.Wait()
			Eventually(session).Should(gbytes.Say("test\\-container\\-1"))
		})

		It("calls dkr c down", func() {
			session = callDkrC("down")
			session.Wait(60000)
			Eventually(session).Should(gbytes.Say("Calling docker-compose"))
			Eventually(session).Should(gbytes.Say("down"))
		})

		It("checks if dkr c down worked", func() {
			session = runMain(mainPath, []string{"ps"})
			session.Wait()
			Consistently(session).ShouldNot(gbytes.Say("test\\-container\\-1"))
		})
	})

	Context("Aliases", func() {
		var spinup = func() *gexec.Session {
			// Spinup
			return callDkrC("up")
		}

		It("spins up", func() {
			session = spinup()
			session.Wait(60000)
		})

		It("checks is spin up worked", func() {
			session = runMain(mainPath, []string{"ps"})
			session.Wait()
			Eventually(session).Should(gbytes.Say("test-container-1"))
		})

		It("calls killall", func() {
			session = runMain(mainPath, []string{"killall"})
			session.Wait()
		})

		It("checks if killall worked", func() {
			session = runMain(mainPath, []string{"ps"})
			session.Wait()
			Consistently(session).ShouldNot(gbytes.Say("test-container-1"))
		})

		It("spins up", func() {
			session = spinup()
			session.Wait(60000)
		})

		It("checks is spin up worked", func() {
			session = runMain(mainPath, []string{"ps"})
			session.Wait()
			Eventually(session).Should(gbytes.Say("test-container-1"))
		})

		It("calls cleaup", func() {
			session = runMain(mainPath, []string{"cleanup"})
			session.Wait(60000)
			Eventually(session).Should(gbytes.Say("Untagged"))
			Eventually(session).Should(gbytes.Say("Deleted"))
		})

		It("spins up", func() {
			session = spinup()
			session.Wait(60000)
		})

		It("checks is spin up worked", func() {
			session = runMain(mainPath, []string{"ps"})
			session.Wait()
			Eventually(session).Should(gbytes.Say("test-container-1"))
		})

		It("callas nuke", func() {
			session = runMain(mainPath, []string{"nuke"})
			session.Wait(60000)
			Eventually(session).Should(gbytes.Say("reclaimed space"))
		})
	})
})

func buildMain() string {
	mainApp, err := gexec.Build("main.go")
	Expect(err).NotTo(HaveOccurred())

	return mainApp
}

func runMain(path string, args []string) *gexec.Session {
	cmd := exec.Command(path, args...)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}
