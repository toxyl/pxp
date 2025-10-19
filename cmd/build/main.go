package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/toxyl/flo"
)

func printBlue(message string) {
	fmt.Printf("\033[1;34m%s\033[0m\n", message)
}

func printYellow(message string) {
	fmt.Printf("\033[1;33m%s\033[0m\n", message)
}

func printGreen(message string) {
	fmt.Printf("\033[1;32m%s\033[0m\n", message)
}

func printRed(message string) {
	fmt.Printf("\033[1;31m%s\033[0m\n", message)
}

func dieOnError(err error, msg string) {
	if err != nil {
		printRed(fmt.Sprintf("%s: %v", msg, err))
		os.Exit(1)
	}
}

func buildGoDSL(src *flo.DirObj, bin *flo.FileObj) error {
	printBlue("Building GoDSL...")

	// Change to GoDSL source directory and build
	cmd := exec.Command("go", "build", "-o", bin.Path(), "-buildvcs=false", ".")
	cmd.Dir = src.Path()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build GoDSL: %w", err)
	}

	// // Move the binary to current directory
	// sourcePath := src.Path()

	// if err := flo.File(sourcePath).Copy(bin.Path()); err != nil {
	// 	return fmt.Errorf("failed to copy GoDSL binary: %w", err)
	// }

	// // Remove the source file
	// if err := flo.File(sourcePath).Remove(); err != nil {
	// 	return fmt.Errorf("failed to remove source GoDSL binary: %w", err)
	// }

	// time.Sleep(5 * time.Second)
	return nil
}

func generateDSL(src *flo.DirObj, bin *flo.FileObj) error {
	printBlue("Generating DSL...")

	// Run the go-dsl binary with the specified parameters
	cmd := exec.Command(bin.Path(),
		"pxp",
		"PixelPipeline Script",
		"A scripting language for defining image processing pipelines.",
		"1.0.0",
		"pxp",
		src.Path())

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate DSL: %w", err)
	}

	// Run go mod tidy
	cmd = exec.Command("go", "mod", "tidy")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run go mod tidy: %w", err)
	}

	time.Sleep(5 * time.Second)
	return nil
}

func main() {
	/////////////////////////////////////////////////////////////////////////////////////////
	printYellow("PixelPipeline Builder")
	/////////////////////////////////////////////////////////////////////////////////////////
	var (
		buildDir        = flag.String("build-dir", "/tmp/build", "Build directory")
		appSlug         = flag.String("slug", "pxp", "Application slug")
		appName         = flag.String("name", "PixelPipeline Studio", "Application name")
		arch            = flag.String("arch", "amd64", "Architecture")
		goDslSourcePath = flag.String("godslsrc", "../go-dsl/app", "Path to go-dsl source directory")
	)
	flag.Parse()

	// Validate required flags
	if *buildDir == "" || *appName == "" || *appSlug == "" || *arch == "" {
		printRed("Error: Missing required flags")
		flag.Usage()
		os.Exit(1)
	}

	/////////////////////////////////////////////////////////////////////////////////////////
	printBlue("Preparing build...")
	/////////////////////////////////////////////////////////////////////////////////////////

	var (
		dSrc         = flo.Dir(".")                     // /src/pxp/
		dSrcBin      = dSrc.Dir("bin")                  // /src/pxp/bin/
		dSrcLang     = dSrc.Dir("language")             // /src/pxp/language/
		dBuild       = flo.Dir("/tmp").Dir(*appSlug)    // /tmp/pxp/
		dApp         = dBuild.Dir("app")                // /tmp/pxp/app/
		dAppFrontend = dApp.Dir("frontend")             // /tmp/pxp/app/frontend/
		dAppBuildBin = dApp.Dir(".build").Dir("bin")    // /tmp/pxp/app/.build/bin/
		dNodeModules = dAppFrontend.Dir("node_modules") // /tmp/pxp/app/frontend/node_modules
		fGoDSLBin    = dSrcBin.File("go-dsl")
		fBinSrc      *flo.FileObj
		fBinDst      *flo.FileObj
	)

	/////////////////////////////////////////////////////////////////////////////////////////
	printBlue("Creating language...")
	/////////////////////////////////////////////////////////////////////////////////////////

	dSrcLang.Each(func(f *flo.FileObj) {
		if matched, _ := filepath.Match("dsl_*", f.Name()); matched {
			if err := f.Remove(); err != nil {
				fmt.Printf("Warning: Failed to remove dsl file %s: %v\n", f.Path(), err)
			}
		}
	}, nil)

	// Remove template_* files using recursive processing
	dSrcLang.Each(func(f *flo.FileObj) {
		if matched, _ := filepath.Match("template_*", f.Name()); matched {
			if err := f.Remove(); err != nil {
				fmt.Printf("Warning: Failed to remove template file %s: %v\n", f.Path(), err)
			}
		}
	}, nil)

	dieOnError(buildGoDSL(flo.Dir(*goDslSourcePath), fGoDSLBin), "Failed to build GoDSL")
	dieOnError(generateDSL(dSrcLang, fGoDSLBin), "Failed to generate DSL")

	for _, osName := range []string{"linux", "windows"} {
		/////////////////////////////////////////////////////////////////////////////////////////
		printBlue(fmt.Sprintf("%s: %s", osName, "Creating directories..."))
		/////////////////////////////////////////////////////////////////////////////////////////

		time.Sleep(5 * time.Second)
		dieOnError(dBuild.Mkdir(0755), "Failed to create build directory")
		dieOnError(dSrcBin.Mkdir(0755), "Failed to create source bin directory")
		dieOnError(dAppBuildBin.Mkdir(0755), "Failed to create app build bin directory")
		dieOnError(dNodeModules.Mkdir(0755), "Failed to create node_modules directory")
		defer dBuild.Remove() // remove once we're done

		/////////////////////////////////////////////////////////////////////////////////////////
		printBlue(fmt.Sprintf("%s: %s", osName, "Copying sources..."))
		/////////////////////////////////////////////////////////////////////////////////////////

		time.Sleep(5 * time.Second)
		dSrc.Each(func(f *flo.FileObj) {
			prel, err := filepath.Rel(dSrc.Path(), f.Path())
			if strings.Contains(prel, "node_modules") {
				return // don't copy node_modules stuff, takes long and we'd remove it anyway
			}
			if strings.HasPrefix(prel, ".git") {
				return // don't copy git stuff, we don't need it
			}
			if strings.HasPrefix(prel, "test_") {
				return // don't copy test stuff, we don't need it
			}
			dieOnError(err, "Coud not make relative path")
			dieOnError(f.Copy(dBuild.File(prel).Path()), "Failed to copy sources")
		}, nil)

		/////////////////////////////////////////////////////////////////////////////////////////
		printBlue(fmt.Sprintf("%s: %s", osName, "Building desktop app..."))
		/////////////////////////////////////////////////////////////////////////////////////////

		time.Sleep(5 * time.Second)
		dieOnError(os.Chdir(dAppFrontend.Path()), "Could not cd into frontend dir")
		cmdNpm := exec.Command("npm", "install")
		cmdNpm.Stdout = nil
		cmdNpm.Stderr = os.Stderr
		dieOnError(cmdNpm.Run(), "NPM install failed")

		time.Sleep(5 * time.Second)
		dieOnError(os.Chdir(dApp.Path()), "Could not cd into source dir")
		os.Setenv("QT_QPA_PLATFORM", "offscreen")
		os.Setenv("QT_OPENGL", "software")
		os.Setenv("QTWEBENGINE_DISABLE_SANDBOX", "1")
		cmdWails := exec.Command("wails", "build", "-clean", "-trimpath", "-platform", fmt.Sprintf("%s/%s", osName, *arch), "-v", "0")
		if osName == "windows" {
			cmdWails.Args = append(cmdWails.Args, "-nsis", "-tags", "\"webkit2_41\"")
		} else {
			cmdWails.Args = append(cmdWails.Args, "-tags", "webkit2_41")
		}
		cmdWails.Args = append(cmdWails.Args, "-o", *appSlug)
		cmdWails.Stdout = os.Stdout
		cmdWails.Stderr = os.Stderr
		dieOnError(cmdWails.Run(), "Wails build failed")

		/////////////////////////////////////////////////////////////////////////////////////////
		printBlue(fmt.Sprintf("%s: %s", osName, "Copying result to source dir..."))
		/////////////////////////////////////////////////////////////////////////////////////////

		time.Sleep(5 * time.Second)

		if osName == "windows" {
			fBinSrc = dAppBuildBin.File(fmt.Sprintf("%s-%s-installer.exe", *appName, *arch))
			fBinDst = dSrcBin.File(fmt.Sprintf("%s-%s-%s-installer.exe", *appSlug, osName, *arch))
		} else {
			fBinSrc = dAppBuildBin.File(*appSlug)
			fBinDst = dSrcBin.File(fmt.Sprintf("%s-%s-%s", *appSlug, osName, *arch))
		}

		dieOnError(fBinSrc.Copy(fBinDst.Path()), "Failed to copy result")
	}

	/////////////////////////////////////////////////////////////////////////////////////////
	printGreen("Build completed successfully!")
	/////////////////////////////////////////////////////////////////////////////////////////
}
